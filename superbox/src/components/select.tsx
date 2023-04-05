import {Select} from "antd";

export interface SelectOption {
    label: string,
    value: string
}

export default function SelectInput(options: SelectOption[]) {
    return (
        <Select
            showSearch
            style={{width: 160}}
            placeholder="Select a person"
            optionFilterProp="children"
            // onChange={onChange}
            // onSearch={onSearch}
            filterOption={(input, option) =>
                (option?.label ?? '').toLocaleString().toLowerCase().includes(input.toLowerCase())
            }
            options={options}
        />
    );
}