import {useState} from "react";
import {Button} from 'antd';

import CookieModal from "./login";


export default function MusicIndex() {

    const [cookie, setCookie] = useState("");
    const [cookieState, setCookieState] = useState(false);

    const setCookies = () => {
        setCookieState(true)
        console.log(cookie)
    }

    const setCookieOnOk = (e: React.MouseEvent) => {
        setCookieState(false)
    }
    const setCookieOnCancel = (e: React.MouseEvent) => {
        setCookieState(false)
    }

    return (
        <div>
            <Button type="primary" onClick={setCookies}>
                Vertically centered modal dialog
            </Button>
            {/* 需要登录的时候调用 */}
            <CookieModal state={cookieState} onOk={setCookieOnOk} onCancel={setCookieOnCancel}
                         cookie={cookie} setCookie={setCookie}/>
        </div>
    )
};