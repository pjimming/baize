import store from "@/store";
import axios from "axios";

const baizeURL = "http://localhost:8888/baize/v1";

export const fetchModulePath = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/module/path`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setModulePath', response.data);

            await fetchPackages(queryParams);
            await fetchModuleName({
                modulePath: response.data.modulePath,
            });

            return response.data.modulePath;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        throw new Error(`${error.message}`);
    }
};

export const fetchPackages = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/packages`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setPackages', response.data);
            return response.data;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        throw new Error(`${error.message}`);
    }
};

export const fetchModuleName = async (queryParams) => {
    try {
        const response = await axios.get(`${baizeURL}/local/module/name`, {
            params: queryParams,
        });

        if (response.status === 200) {
            store.commit('setModuleName', response.data);
            return response.data;
        } else {
            throw new Error(`Request fail with status: ${response.status}`);
        }
    } catch (error) {
        throw new Error(`${error.message}`);
    }
};