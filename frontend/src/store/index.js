import { createStore } from 'vuex'

export default createStore({
  state: {
    modulePath: "",
    dir: "",
    projectInfo: {
      modulePath: "",
      moduleName: "",
      packageCount: 0,
      packageList: [],
    },
  },
  getters: {
  },
  mutations: {
    setModulePath(state, data) {
      state.modulePath = data;
    },
    setDir(state, data) {
      state.dir = data;
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
