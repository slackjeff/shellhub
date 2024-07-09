import { Module } from "vuex";

import * as apiDevice from "../api/devices";
import * as apiBilling from "../api/billing";
import { IDevice } from "@/interfaces/IDevice";
import { State } from "..";

export interface DevicesState {
  devices: Array<IDevice>;
  quickConnectionList: Array<IDevice>;
  device: IDevice;
  numberDevices: number;
  page: number;
  perPage: number;
  filter: undefined | string;
  status: "accepted" | "rejected" | "pending" | "unused";
  sortStatusField: undefined | string;
  sortStatusString: "asc" | "desc" | "";
  deviceChooserStatus: boolean;
  devicesForUserToChoose: Array<IDevice>;
  numberdevicesForUserToChoose: number;
  devicesSelected: Array<IDevice>;
  deviceName: string,
  }

export const devices: Module<DevicesState, State> = {
  namespaced: true,
  state: {
    devices: [],
    quickConnectionList: [],
    device: {} as IDevice,
    numberDevices: 0,
    page: 1,
    perPage: 10,
    filter: "",
    status: "accepted",
    sortStatusField: undefined,
    sortStatusString: "asc",
    deviceChooserStatus: false,
    devicesForUserToChoose: [],
    numberdevicesForUserToChoose: 0,
    devicesSelected: [],
    deviceName: "",
  },

  getters: {
    list: (state) => state.devices,
    listQuickConnection: (state) => state.quickConnectionList,
    get: (state) => state.device,
    getName: (state) => state.device.name,
    getNumberDevices: (state) => state.numberDevices,
    getPage: (state) => state.page,
    getPerPage: (state) => state.perPage,
    getFilter: (state) => state.filter,
    getStatus: (state) => state.status,
    getSortStatusField: (state) => state.sortStatusField,
    getSortStatusString: (state) => state.sortStatusString,
    getFirstPending: (state) => state.device,
    getDeviceChooserStatus: (state) => state.deviceChooserStatus,
    getDevicesForUserToChoose: (state) => state.devicesForUserToChoose,
    getNumberForUserToChoose: (state) => state.numberdevicesForUserToChoose,
    getDevicesSelected: (state) => state.devicesSelected,
    getDeviceToBeRenamed: (state) => state.deviceName,
  },

  mutations: {
    setDevices: (state, res, committable = true) => {
      if (committable) {
        state.devices = res.data;
      }
      state.numberDevices = parseInt(res.headers["x-total-count"], 10);
    },

    setTotalCount: (state, res) => {
      state.numberDevices = parseInt(res.headers["x-total-count"], 10);
    },

    setQuickDevices: (state, res) => {
      state.quickConnectionList = res.data;
    },

    clearQuickDevices: (state) => {
      state.quickConnectionList = [];
    },

    removeDevice: (state, uid) => {
      state.devices.splice(
        state.devices.findIndex((d) => d.uid === uid),
        1,
      );
    },

    renameDevice: (state, data) => {
      const { device } = state;
      device.name = data.name;
      state.device = device;
    },

    setDevice: (state, data) => {
      state.device = data;
    },

    setPagePerpageFilter: (state, data) => {
      state.page = data.page;
      state.perPage = data.perPage;
      state.filter = data.filter;
      state.status = data.status;
      state.sortStatusField = data.sortStatusField;
      state.sortStatusString = data.sortStatusString;
    },

    setFilter: (state, filter) => {
      state.filter = filter;
    },

    setDeviceChooserStatus: (state, status) => {
      state.deviceChooserStatus = status;
    },

    setDevicesForUserToChoose: (state, res) => {
      state.devicesForUserToChoose = res.data;
      state.numberdevicesForUserToChoose = parseInt(res.headers["x-total-count"] ?? 1, 10);
    },

    setDevicesSelected: (state, data) => {
      state.devicesSelected = data;
    },

    setSortStatusDevice: (state, data) => {
      state.sortStatusField = data.sortStatusField;
      state.sortStatusString = data.sortStatusString;
    },

    clearListDevices: (state) => {
      state.devices = [];
      state.numberDevices = 0;
    },

    clearObjectDevice: (state) => {
      state.device = {} as IDevice;
    },

    clearListDevicesForUserToChoose: (state) => {
      state.devicesForUserToChoose = [];
      state.numberdevicesForUserToChoose = 0;
    },

    updateDeviceToBeRenamed(state, device) {
      state.deviceName = device;
    },
  },

  actions: {
    fetch: async ({ commit }, data) => {
      try {
        const res = await apiDevice.fetchDevices(
          data.page,
          data.perPage,
          data.filter,
          data.status,
          data.sortStatusField,
          data.sortStatusString,
        );
        if (res.data.length) {
          commit("setDevices", res, data.commitable);
          commit("setPagePerpageFilter", data);
          return res;
        }

        commit("clearListDevices");
        return false;
      } catch (error) {
        commit("clearListDevices");
        throw error;
      }
    },

    remove: async (context, uid) => {
      await apiDevice.removeDevice(uid);
    },

    rename: async (context, data) => {
      await apiDevice.renameDevice(data);
      context.commit("renameDevice", data);
    },

    get: async (context, uid) => {
      try {
        const res = await apiDevice.getDevice(uid);
        context.commit("setDevice", res.data);
      } catch (error) {
        context.commit("clearObjectDevice");
        throw error;
      }
    },

    accept: async (context, uid) => {
      try {
        await apiDevice.acceptDevice(uid);
      } catch (error) {
        console.error(error);
        throw error;
      }
    },

    reject: async (context, uid) => {
      try {
        await apiDevice.rejectDevice(uid);
      } catch (error) {
        console.error(error);
        throw error;
      }
    },

    setFilter: async (context, filter) => {
      context.commit("setFilter", filter);
    },

    refresh: async ({ commit, state }) => {
      try {
        const res = await apiDevice.fetchDevices(
          state.page,
          state.perPage,
          state.filter,
          state.status,
          state.sortStatusField,
          state.sortStatusString,
        );
        commit("setDevices", res);
      } catch (error) {
        commit("clearListDevices");
        throw error;
      }
    },

    async search({ commit, state }, data) {
      try {
        const res = await apiDevice.fetchDevices(
          data.page,
          data.perPage,
          data.filter,
          state.status,
          state.sortStatusField,
          state.sortStatusString,
        );
        commit("setDevices", res);
        commit("setDevicesForUserToChoose", res);
        commit("setFilter", data.filter);
      } catch (error) {
        commit("clearListDevices");
        throw error;
      }
    },

    fetchQuickDevices: async ({ commit }, data) => {
      try {
        const res = await apiDevice.fetchDevices(
          data.page,
          data.perPage,
          data.filter,
          data.status,
          data.sortStatusField,
          data.sortStatusString,
        );
        if (res.data.length) {
          commit("setQuickDevices", res);
          return res;
        }

        commit("clearQuickDevices");
        return false;
      } catch (error) {
        commit("clearQuickDevices");
        throw error;
      }
    },

    async searchQuickConnection({ commit, state }, data) {
      try {
        const res = await apiDevice.fetchDevices(
          data.page,
          data.perPage,
          data.filter,
          state.status,
          state.sortStatusField,
          state.sortStatusString,
        );
        commit("setQuickDevices", res);
      } catch (error) {
        commit("clearQuickDevices");
        throw error;
      }
    },

    setFirstPending: async (context) => {
      try {
        const res = await apiDevice.fetchDevices(
          1,
          1,
          undefined,
          "pending",
          undefined,
          "",
        );
        context.commit("setDevice", res.data[0]);
      } catch (error) {
        context.commit("clearObjectDevice");
        throw error;
      }
    },

    setDeviceChooserStatus: async (context, status) => {
      context.commit("setDeviceChooserStatus", status);
    },

    setDevicesForUserToChoose: async (context, data) => {
      try {
        const res = await apiDevice.fetchDevices(
          data.page,
          data.perPage,
          data.filter,
          data.status,
          data.sortStatusField,
          data.sortStatusString,
        );
        res.data = res.data.map((item) => {
          const newItem = item;
          delete newItem.last_seen;
          return newItem;
        });
        if (res.data.length) {
          context.commit("setDevicesForUserToChoose", res);
          context.commit("setPagePerpageFilter", data);
          return res;
        }

        return false;
      } catch (error) {
        context.commit("clearListDevicesForUserToChoose");
        throw error;
      }
    },

    setDevicesSelected: (context, data) => {
      context.commit("setDevicesSelected", data);
    },

    async setSortStatus({ commit }, data) {
      commit("setSortStatusDevice", data);
    },

    postDevicesChooser: async (context, data) => {
      try {
        await apiBilling.postDevicesChooser(data);
      } catch (error) {
        console.error(error);
        throw error;
      }
    },

    getDevicesMostUsed: async (context) => {
      try {
        const res = await apiBilling.getDevicesMostUsed();
        context.commit("setDevicesForUserToChoose", res);
      } catch (error) {
        context.commit("clearListDevicesForUserToChoose");
        throw error;
      }
    },

    resetListDevices: async (context) => {
      context.commit("clearListDevices");
    },

    updateDeviceTag: async (context, data) => {
      try {
        await apiDevice.updateDeviceTag(data);
      } catch (error) {
        console.error(error);
        throw error;
      }
    },

    setDeviceToBeRenamed(context, device) {
      context.commit("updateDeviceToBeRenamed", device);
    },
  },
};
