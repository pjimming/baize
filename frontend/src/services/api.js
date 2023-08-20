import store from "@/store";
import axios from "axios";

const baizeURL = "http://localhost:8888/baize/v1";

export const fetchModulePath = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/module`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setModulePath', response.data.modulePath);
            return response.data.modulePath;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        throw new Error(`${error.message}`);
    }
};

export const fetchProjectInfo = async (queryParams) => {
    console.log("fetchProjectInfo", queryParams);
    try {
        const response = await axios.get(`${baizeURL}/local/info`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setProjectInfo', response.data);
            return response.data;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        throw new Error(`${error.message}`);
    }
};

