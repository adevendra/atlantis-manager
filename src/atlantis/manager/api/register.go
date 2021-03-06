/* Copyright 2014 Ooyala, Inc. All rights reserved.
 *
 * This file is licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is
 * distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package api

import (
	. "atlantis/common"
	. "atlantis/manager/rpc/types"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ListRouters(w http.ResponseWriter, r *http.Request) {
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	internal, err := strconv.ParseBool(r.FormValue("Internal"))
	if err != nil {
		fmt.Fprintf(w, "%s", Output(map[string]interface{}{}, err))
		return
	}
	arg := ManagerListRoutersArg{ManagerAuthArg: auth, Internal: internal}
	var reply ManagerListRoutersReply
	err = manager.ListRouters(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Routers": reply.Routers, "Status": reply.Status}, err))
}

func RegisterRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	internal, err := strconv.ParseBool(r.FormValue("Internal"))
	if err != nil {
		fmt.Fprintf(w, "%s", Output(map[string]interface{}{}, err))
		return
	}
	arg := ManagerRegisterRouterArg{
		ManagerAuthArg: auth,
		Internal:       internal,
		Zone:           vars["Zone"],
		Host:           vars["Host"],
		IP:             r.FormValue("IP"),
	}
	var reply AsyncReply
	err = manager.RegisterRouter(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"ID": reply.ID}, err))
}

func UnregisterRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	internal, err := strconv.ParseBool(r.FormValue("Internal"))
	if err != nil {
		fmt.Fprintf(w, "%s", Output(map[string]interface{}{}, err))
		return
	}
	arg := ManagerRegisterRouterArg{
		ManagerAuthArg: auth,
		Internal:       internal,
		Zone:           vars["Zone"],
		Host:           vars["Host"],
	}
	var reply AsyncReply
	err = manager.UnregisterRouter(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"ID": reply.ID}, err))
}

func GetRouter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	internal, err := strconv.ParseBool(r.FormValue("Internal"))
	if err != nil {
		fmt.Fprintf(w, "%s", Output(map[string]interface{}{}, err))
		return
	}
	arg := ManagerGetRouterArg{
		ManagerAuthArg: auth,
		Internal:       internal,
		Zone:           vars["Zone"],
		Host:           vars["Host"],
	}
	var reply ManagerGetRouterReply
	err = manager.GetRouter(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Router": reply.Router}, err))
}

func ListRegisteredApps(w http.ResponseWriter, r *http.Request) {
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	authorizedOnly, _ := strconv.ParseBool(r.FormValue("AuthorizedOnly"))
	if authorizedOnly {
		arg := ManagerListRegisteredAppsArg{auth}
		var reply ManagerListRegisteredAppsReply
		err := manager.ListAuthorizedRegisteredApps(arg, &reply)
		fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Apps": reply.Apps, "Status": reply.Status}, err))
	} else {
		arg := ManagerListRegisteredAppsArg{auth}
		var reply ManagerListRegisteredAppsReply
		err := manager.ListRegisteredApps(arg, &reply)
		fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Apps": reply.Apps, "Status": reply.Status}, err))
	}
}

func RegisterApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	nonAtlantis, _ := strconv.ParseBool(r.FormValue("NonAtlantis"))
	internal, _ := strconv.ParseBool(r.FormValue("Internal"))
	arg := ManagerRegisterAppArg{
		ManagerAuthArg: auth,
		NonAtlantis:    nonAtlantis,
		Internal:       internal,
		Name:           vars["App"],
		Repo:           r.FormValue("Repo"),
		Root:           r.FormValue("Root"),
		Email:          r.FormValue("Email"),
	}
	var reply ManagerRegisterAppReply
	err := manager.RegisterApp(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status}, err))
}

func UpdateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	nonAtlantis, _ := strconv.ParseBool(r.FormValue("NonAtlantis"))
	internal, _ := strconv.ParseBool(r.FormValue("Internal"))
	arg := ManagerRegisterAppArg{
		ManagerAuthArg: auth,
		NonAtlantis:    nonAtlantis,
		Internal:       internal,
		Name:           vars["App"],
		Repo:           r.FormValue("Repo"),
		Root:           r.FormValue("Root"),
		Email:          r.FormValue("Email"),
	}
	var reply ManagerRegisterAppReply
	err := manager.UpdateApp(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status}, err))
}

func UnregisterApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRegisterAppArg{ManagerAuthArg: auth, Name: vars["App"]}
	var reply ManagerRegisterAppReply
	err := manager.UnregisterApp(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status}, err))
}

func GetApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerGetAppArg{ManagerAuthArg: auth, Name: vars["App"]}
	var reply ManagerGetAppReply
	err := manager.GetApp(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "App": reply.App}, err))
}

func ListSupervisors(w http.ResponseWriter, r *http.Request) {
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerListSupervisorsArg{auth}
	var reply ManagerListSupervisorsReply
	err := manager.ListSupervisors(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Supervisors": reply.Supervisors, "Status": reply.Status}, err))
}

func RegisterSupervisor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRegisterSupervisorArg{auth, vars["Host"]}
	var reply AsyncReply
	err := manager.RegisterSupervisor(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"ID": reply.ID}, err))
}

func UnregisterSupervisor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRegisterSupervisorArg{auth, vars["Host"]}
	var reply AsyncReply
	err := manager.UnregisterSupervisor(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"ID": reply.ID}, err))
}

func ListManagers(w http.ResponseWriter, r *http.Request) {
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerListManagersArg{auth}
	var reply ManagerListManagersReply
	err := manager.ListManagers(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Managers": reply.Managers, "Status": reply.Status}, err))
}

func RegisterManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRegisterManagerArg{
		ManagerAuthArg: auth,
		Host:           vars["Host"],
		Region:         vars["Region"],
		ManagerCName:   r.FormValue("ManagerCName"),
		RegistryCName:  r.FormValue("RegistryCName"),
	}
	var reply AsyncReply
	err := manager.RegisterManager(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"ID": reply.ID}, err))
}

func UnregisterManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRegisterManagerArg{ManagerAuthArg: auth, Host: vars["Host"], Region: vars["Region"]}
	var reply AsyncReply
	err := manager.UnregisterManager(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"ID": reply.ID}, err))
}

func GetManager(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerGetManagerArg{
		ManagerAuthArg: auth,
		Region:         vars["Region"],
		Host:           vars["Host"],
	}
	var reply ManagerGetManagerReply
	err := manager.GetManager(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Manager": reply.Manager}, err))
}

func GetSelf(w http.ResponseWriter, r *http.Request) {
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerGetSelfArg{ManagerAuthArg: auth}
	var reply ManagerGetManagerReply
	err := manager.GetSelf(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Manager": reply.Manager}, err))
}

func AddRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRoleArg{
		ManagerAuthArg: auth,
		Host:           vars["Host"],
		Region:         vars["Region"],
		Role:           vars["Role"],
	}
	var reply ManagerRoleReply
	err := manager.AddRole(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Manager": reply.Manager}, err))
}

func RemoveRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRoleArg{
		ManagerAuthArg: auth,
		Host:           vars["Host"],
		Region:         vars["Region"],
		Role:           vars["Role"],
	}
	var reply ManagerRoleReply
	err := manager.RemoveRole(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Manager": reply.Manager}, err))
}

func AddRoleType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRoleArg{
		ManagerAuthArg: auth,
		Host:           vars["Host"],
		Region:         vars["Region"],
		Role:           vars["Role"],
		Type:           vars["Type"],
	}
	var reply ManagerRoleReply
	err := manager.AddRole(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Manager": reply.Manager}, err))
}

func RemoveRoleType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	auth := ManagerAuthArg{r.FormValue("User"), "", r.FormValue("Secret")}
	arg := ManagerRoleArg{
		ManagerAuthArg: auth,
		Host:           vars["Host"],
		Region:         vars["Region"],
		Role:           vars["Role"],
		Type:           vars["Type"],
	}
	var reply ManagerRoleReply
	err := manager.RemoveRole(arg, &reply)
	fmt.Fprintf(w, "%s", Output(map[string]interface{}{"Status": reply.Status, "Manager": reply.Manager}, err))
}
