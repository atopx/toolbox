import {Input, Modal} from 'antd';

interface IProps {
    cookie: string,
    setCookie: (cookie: string) => void;
    state: boolean;
    onOk?: (e: React.MouseEvent<HTMLButtonElement>) => void;
    onCancel?: (e: React.MouseEvent<HTMLButtonElement>) => void;
}

const CookieModal = (props: IProps) => {
    const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        console.log(event.target.value);
        props.setCookie(event.target.value)
    }
    return (
        <Modal
            title="设置QQ音乐Cookie"
            centered
            open={props.state}
            onOk={props.onOk}
            onCancel={props.onCancel}
        >
            <Input value={props.cookie} onChange={onChange}/>
        </Modal>
    );
}


export default CookieModal;