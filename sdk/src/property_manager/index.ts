import axios, { AxiosInstance } from 'axios';

export class PropertyManager {
    private url: string;
    private client: AxiosInstance;

    constructor(url: string) {
        this.url = url;
        this.client = axios.create({
            baseURL: url,
            timeout: 2000
        });
    }

    async getProperty(key: string): Promise<string> {
        const resp = await this.client.get(`/key=${key}`)
        console.log(resp);
        return '';
    }

    async setProperty(key: string, value: string): Promise<string> {
        // TODO
        return "";
    }
}