import { createStore } from 'vuex';

const initialState = {
  modulePath: "",
  dir: "",
  projectInfo: {
    modulePath: "",
    moduleName: "",
    moduleVersion: "",
    otherPkgList: [],
    otherPkgCount: 0,
    projectPkgList: [],
    projectPkgCount: 0,
    goFileList: [],
    goFileCount: 0,
  },
};

export default createStore({
  state: { ...initialState },
  getters: {
  },
  mutations: {
    resetState(state) {
      Object.assign(state, initialState);
    },
    setModuleInfo(state, data) {
      state.projectInfo.modulePath = data.modulePath;
      state.projectInfo.moduleName = data.moduleName;
      state.projectInfo.moduleVersion = "go " + data.moduleVersion;
    },
    setDir(state, data) {
      state.dir = data;
    },
    setPackages(state, data) {
      state.projectInfo.otherPkgList = data.otherPkgList;
      state.projectInfo.otherPkgCount = data.otherPkgCount;
      state.projectInfo.projectPkgList = data.projectPkgList;
      state.projectInfo.projectPkgCount = data.projectPkgCount;
    },
    setGoFiles(state, data) {
      state.projectInfo.goFileList = data.goFileList;
      state.projectInfo.goFileCount = data.goFileCount;
    },
    setModuleName(state, data) {
      state.projectInfo.moduleName = data.moduleName;
    },
    setProjectInfo(state, data) {
      state.projectInfo.modulePath = state.modulePath;
      state.projectInfo.moduleName = data.moduleName;
      state.projectInfo.packageCount = data.packageCount;
      state.projectInfo.packageList = data.packageList;
    }
  },
  actions: {
  },
  modules: {
  }
})
