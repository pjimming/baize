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

            return response.data.modulePath;
        } else {
            throw new Error(`Request fail with status: ${response}`);
        }
    } catch (error) {
        console.log("errpr:", error);
        ElMessage({
            showClose: true,
            message: `${error.response.data.desc}`,
            type: 'error',
        })
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
        })
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
        })
    }
}