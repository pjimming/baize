import store from "@/store";
import axios from "axios";
import { ElMessage } from "element-plus";

const baizeURL = "http://localhost:8888/baize/v1";

export const fetchModuleInfo = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/module`, {
            params: queryParams,
        });
        if (response.status === 200) {

            store.commit('setModuleInfo', response.data.result);

            await fetchPackages(queryParams);
            await fetchGoFiles(queryParams);
            await fetchLocalGraph(queryParams);

            return response.data.modulePath;
        } else {
            throw new Error(`Request fail with status: ${response}`);
        }
    } catch (error) {
        if (error.code === "ERR_NETWORK") {
            ElMessage({
                showClose: true,
                message: "网络错误",
                type: 'error',
            });
        } else {
            ElMessage({
                showClose: true,
                message: `${error.response.data.desc}`,
                type: 'error',
            });
        }
    }
};

export const fetchPackages = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/packages`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setPackages', response.data.result);
            return response.data;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        ElMessage({
            showClose: true,
            message: `${error.response.data.desc}`,
            type: 'error',
        });
    }
};

export const fetchGoFiles = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/file`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setGoFiles', response.data.result);
            return response.data;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        ElMessage({
            showClose: true,
            message: `${error.response.data.desc}`,
            type: 'error',
        });
    }
};

export const fetchLocalGraph = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/graph`, {
            params: queryParams,
        });

        if (response.status === 200) {
            console.log(response, "local graph");
            store.commit('setGraph', response.data.result);
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        ElMessage({
            showClose: true,
            message: `${error.response.data.desc}`,
            type: 'error',
        });
    }
};