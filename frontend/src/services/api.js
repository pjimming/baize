import store from "@/store";
import axios from "axios";

const baizeURL = "http://localhost:8888/baize/v1";

export const fetchModuleInfo = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/module`, {
            params: queryParams,
        });
        console.log(response);
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
        throw new Error(`${error.response.data.desc}`);
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
        throw new Error(`${error.message}`);
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
        throw new Error(`${error.message}`);
    }
}