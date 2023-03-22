import { Page, ResponseData } from '~/common/types';
import http from '~/common/request';


type ComputerItem = {
    id: number,
    name: string,
    address: string,
    lan_hostname: string,
    wan_hostname: string,
    username: string,
    password: string,
    create_time: string,
    update_time: string,
}


type ComputerData = {
    page: Page,
    search: Array<ComputerItem>,
}



class Computer {
    static async list(params: { page: Page, filter: any }): Promise<ResponseData<ComputerData>> {
        return await http.post("/computer/search", params);
    }

    static async create(params: {
        name: string,
        username: string,
        password: string,
        address: string,
        lan_hostname: string,
        wan_hostname: string,
        default_ports: Array<number>,
    }): Promise<ResponseData<null>> {
        return await http.post("/computer/create", params)
    }

    static async update(params: {
        id: number,
        name: string,
        address: string,
        lan_hostname: string,
        wan_hostname: string,
        username: string,
        password: string,
    }): Promise<ResponseData<null>> {
        return await http.post("/computer/update", params)
    }

    static async delete(params: { id: number }): Promise<ResponseData<null>> {
        return await http.post("/computer/delete", params)
    }

    static async operate(params: { id: number, operate: number }): Promise<ResponseData<null>> {
        return await http.post("/computer/operate", params)
    }
}


export {
    Computer,
    ComputerItem,
}