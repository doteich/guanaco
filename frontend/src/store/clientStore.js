import { defineStore, storeToRefs } from "pinia";
import { AddClient, DisconnectClient, ReconnectClient, GetClients, AppBrowse, ExportBrowseSelection, StartMonitor, StopMonitor, SaveConfigToFile, LoadConfigFromFile } from "../../wailsjs/go/main/App"


export const useClientStore = defineStore("clientStore", {
    state: () => ({
        clients: [],
        toast: {
            severity: "",
            summary: "",
            detail: "",
            life: 3000,
        },
        selectedClient: -1,
        browseResults: [{
            name: "Root",
            nodeId: "i=84",
            type: "Object",
            dataType: "",
            icon: "pi-folder",
            id: "0",
            childs: []
        }],
        monitoredItems: {

        }

    }),
    getters: {
        getClients(state) {
            return state.clients
        },
        getToast(state) {
            return state.toast
        },
        getBrowseResults(state) {

            function getChilds(childs) {
                let nodes = []

                childs.forEach((child) => {
                    let data = {
                        id: child.id,
                        nodeId: child.nodeId,
                        dataType: child.dataType,
                        name: child.name,
                        icon: child.icon,
                        color: child.color,
                        type: child.type,
                        isExpanded: false
                    }

                    child.childs.length > 0 ? data.isExpanded = true : ""

                    nodes.push(data)
                    if (child.childs.length < 1) {
                        return
                    } else {
                        nodes.push(...getChilds(child.childs))
                    }
                })
                return nodes
            }

            let nodes = getChilds(state.browseResults)
            return nodes

        },
        getMonitoredItems(state) {
            return state.monitoredItems
        }
    },
    actions: {
        listen() {
            window.runtime.EventsOn("client-message", (id, event) => {
                switch (event) {
                    case 'disconnect':
                        this.clients.filter(c => c.id == id)[0].status = "disconnected"
                        break;
                    case 'reconnect':
                        this.clients.filter(c => c.id == id)[0].status = "connected"
                        break;
                    case 'keepalive':
                        console.log(id, event)
                        break;
                }
            }),
                window.runtime.EventsOn("node-monitor", (id, event, value, nodeId, ts) => {
                    switch (event) {
                        case "error":
                            break
                        case "update":
                            let monitor = this.monitoredItems[id]
                            if (monitor) {
                                let node = monitor.items.find(n => n.nodeId == nodeId)
                                if (node) {
                                    node.value = value
                                    node.ts = ts
                                }
                            }


                            break
                    }
                })

        },
        async addClient(name, ep, mode, policy, auth, user, password) {
            AddClient(name, ep, mode, policy, auth, user, password)
                .then((data) => {
                    this.clients.push({
                        id: data,
                        name: name,
                        ip: ep,
                        mode,
                        policy,
                        auth,
                        user,
                        password,
                        status: "connected"
                    })
                    if (this.selectedClient == -1) {
                        this.selectClient(this.clients[0].id)
                    }
                })
                .catch((err) => {
                    this.toast = { severity: "error", summary: "Failed to Add OPC UA Client", detail: err, life: 5000 }
                })
        },
        async disconnectClient(id) {
            DisconnectClient(id)

        },
        async reconnect(id) {
            ReconnectClient(id)
        },
        async getActiveConnections() {
            GetClients()
                .then(res => {
                    res.forEach((client) => {
                        this.clients.push({
                            id: client.ClientId,
                            name: client.Name,
                            status: client.Status,
                            ip: client.IP,
                            mode: client.Mode,
                            policy: client.Policy,
                            auth: client.Auth,
                            user: client.User,
                            password: client.Password

                        })
                    })
                    if (this.selectedClient == -1) {

                        this.selectClient(this.clients[0].id)
                    }
                })
        },
        selectClient(id) {

            if (this.selectedClient > -1) {
                if (this.selectedClient != id) {
                    this.browseResults[0].childs = []
                }

                let s = this.clients.find(c => c.id == this.selectedClient)
                s.selected = false
                this.selectedClient = -1


            }

            let c = this.clients.find(c => c.id == id)
            if (c) {
                c.selected = true
                this.selectedClient = id
            }
        },
        Browse(nodeId, index) {
            if (this.selectedClient == -1) {
                this.toast = {
                    severity: "warn",
                    summary: "No Client Selected",
                    detail: "Please select a connected client from the list",
                    life: 3000,
                }
                return
            }
            AppBrowse(this.selectedClient, nodeId)
                .then((res) => {
                    let idc = index.split(".").map(Number)
                    let branch = this.browseResults

                    console.log(res)

                    for (const level of idc) {
                        if (branch[level] && branch[level].childs) {
                            branch = branch[level].childs;
                        } else {
                            return;
                        }
                    }
                    if (branch.length > 0) {

                        branch.length = 0
                        return
                    }

                    res.forEach((node, i) => {

                        let data = {
                            name: node.Name,
                            nodeId: node.NodeId,
                            type: node.Type,
                            dataType: node.DataType,
                            icon: "pi-folder",
                            id: index + "." + i,
                            childs: []
                        }

                        if (node.Type == "NodeClassObject") {
                            data.icon = "pi-folder"
                            data.color = "var(--theme-color-3)"
                        } else {
                            data.icon = "pi-tag"
                            data.color = "rgb(231, 9, 120)"
                        }

                        branch.push(data)
                    })

                    //this.browseResults[0].childs = res
                })
                .catch((err) => {
                    this.toast = {
                        severity: "error",
                        summary: "Browse Error",
                        detail: err,
                        life: 3000,
                    }
                    console.error(err)
                })
        },
        ExportBrowsedNodes(nodes) {

            function getPath(iString, res) {
                let ids = iString.split(".").map(Number)

                let path = ""
                let branch = res


                for (let id of ids) {
                    path = path + branch[id].name + "/"
                    branch = branch[id].childs
                }
                return path

            }

            let exp = []

            nodes.forEach(el => {
                exp.push({
                    name: el.name,
                    nodeId: el.nodeId,
                    dataType: el.dataType,
                    path: getPath(el.id, this.browseResults)
                })
            })

            ExportBrowseSelection(JSON.stringify(exp), this.clients.find(c => c.id == this.selectedClient)?.name)
                .then((res) => {
                    this.toast = {
                        severity: "success",
                        summary: "Exported",
                        detail: "File:" + res + " created",
                        life: 3000,
                    }
                })
                .catch((err) => {
                    this.toast = {
                        severity: "error",
                        summary: "Export Error",
                        detail: err,
                        life: 3000,
                    }
                    console.error(err)
                })
        },
        CreateNodeMonitor(nodes) {
            if (this.monitoredItems.hasOwnProperty(this.selectedClient)) {
                nodes.forEach(node => {
                    this.monitoredItems[this.selectedClient].items.push({
                        nodeId: node.nodeId,
                        value: "",
                        name: node.name,
                        dataType: node.dataType
                    })
                })
            } else {

                let name = this.clients.find(c => c.id == this.selectedClient).name

                this.monitoredItems[this.selectedClient] = {
                    items: [],
                    name: name,
                    isExpanded: false,
                }

                nodes.forEach(node => {
                    this.monitoredItems[this.selectedClient].items.push({
                        nodeId: node.nodeId,
                        value: "",
                        name: node.name,
                        dataType: node.dataType
                    })
                })
            }


            StartMonitor(this.selectedClient, 10, nodes.map(el => el.nodeId))
        },
        stopNodeMonitor(id) {
            
            StopMonitor(Number(id))
            .then(()=> {
                delete(this.monitoredItems[id])
            })
            .catch(err => {
                this.toast = {
                    severity: "error",
                    summary: "Monitor Error",
                    detail: err,
                    life: 3000,
                }
            })
        },

        toggleTable(id, bool) {
            this.monitoredItems[id].isExpanded = bool
        },
        saveConfig() {
            let conf = []
            this.clients.forEach(c => {
                conf.push({
                    name: c.name,
                    ip: c.ip,
                    policy: c.policy,
                    mode: c.mode,
                    auth: c.auth,
                    user: c.user,
                    password: c.password,
                })
            })
            SaveConfigToFile(JSON.stringify(conf))
                .then(res => {
                    this.toast = {
                        severity: "success",
                        summary: "Config Saved",
                        detail: `Saved to: ${res}`,
                        life: 3000,
                    }
                })
                .catch(err => {
                    this.toast = {
                        severity: "error",
                        summary: "Save Error",
                        detail: err,
                        life: 3000,
                    }
                })
        },
        loadConfig() {
            LoadConfigFromFile()
                .then(res => {
                    let j = JSON.parse(res)

                    if (j.length < 1) {
                        this.toast = {
                            severity: "error",
                            summary: "Load Error ",
                            detail: "loaded config invalid or empty",
                            life: 3000,
                        }
                    }
                    j.forEach(c => {
                        this.addClient(c.name, c.ip, c.mode, c.policy, c.auth, c.user, c.password)
                    })
                })
                .catch(err => {
                    this.toast = {
                        severity: "error",
                        summary: "Load Error ",
                        detail: err,
                        life: 3000,
                    }
                })
        }
    }

})


