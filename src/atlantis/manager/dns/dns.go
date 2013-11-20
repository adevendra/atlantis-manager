package dns

import (
	. "atlantis/manager/constant"
	"atlantis/manager/datamodel"
	"atlantis/manager/helper"
	"fmt"
)

var Provider DNSProvider

type DNSProvider interface {
	CreateAliases(string, []Alias) (error, chan error)
	CreateCNames(string, []CName) (error, chan error)
	DeleteRecords(string, ...string) (error, chan error)
	CreateHealthCheck(string) (string, error)
	DeleteHealthCheck(string) error
	Suffix() string
}

type Alias struct {
	Primary  bool
	Alias    string
	Original string
}

func (a *Alias) Id() string {
	return fmt.Sprintf("%t-%s-%s", a.Primary, a.Alias, a.Original)
}

type CName struct {
	Primary       bool
	CName         string
	IP            string
	HealthCheckId string
}

func (c *CName) Id() string {
	return fmt.Sprintf("%t-%s-%s", c.Primary, c.CName, c.IP)
}

func CreateAppAliases(app, sha, env string) error {
	// check if records were created already, if so add sha to list
	zkDNS, err := datamodel.GetDNS(app, env)
	if zkDNS != nil && err == nil {
		zkDNS.Shas[sha] = true
		return zkDNS.Save()
	}
	// for each zone
	zkDNS = datamodel.DNS(app, env)
	zkDNS.Shas[sha] = true
	if Provider == nil {
		return zkDNS.Save()
	}
	zkDNS.RecordIds = make([]string, len(AvailableZones)+1)
	aliases := make([]Alias, len(AvailableZones)+1)
	for i, zone := range AvailableZones {
		aliases[i] = Alias{
			Primary:  true,
			Alias:    helper.GetZoneAppAlias(app, env, zone, Provider.Suffix()),
			Original: helper.GetZoneRouterCName(zone, Provider.Suffix()),
		}
		zkDNS.RecordIds[i] = aliases[i].Id()
	}
	// region-wide entry (for referencing outside of atlantis)
	aliases[len(aliases)-1] = Alias{
		Primary:  true,
		Alias:    helper.GetRegionAppAlias(app, env, Provider.Suffix()),
		Original: helper.GetRegionRouterCName(Provider.Suffix()),
	}
	zkDNS.RecordIds[len(aliases)-1] = aliases[len(aliases)-1].Id()
	err, errChan := Provider.CreateAliases("CREATE_APP "+app+" in "+env, aliases)
	if err != nil {
		return err
	}
	err = <-errChan // wait for change to propagate
	if err != nil {
		return err
	}
	// save records made in router zone path
	return zkDNS.Save()
}

func DeleteAppAliases(app, sha, env string) error {
	// find ids for app+env
	zkDNS, err := datamodel.GetDNS(app, env)
	if err != nil {
		return err
	}
	// remove sha from sha references
	delete(zkDNS.Shas, sha)
	err = zkDNS.Save()
	if err != nil {
		return err
	}
	// if this was *not* the last sha, don't delete anything
	if len(zkDNS.Shas) > 0 {
		return nil
	}
	// delete all the record ids
	err, errChan := Provider.DeleteRecords("DELETE_APP "+app+" in "+env, zkDNS.RecordIds...)
	if err != nil {
		return err
	}
	err = <-errChan // wait for change to propagate
	if err != nil {
		return err
	}
	// remove dns datamodel
	return zkDNS.Delete()
}