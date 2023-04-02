mod curl;
mod json;
mod sql;

use std::collections::HashMap;
use std::str::FromStr;

pub struct ParseError {
    message: &'static str,
}

pub enum Method {
    /// GET 方法请求一个指定资源的表示形式，使用 GET 的请求应该只被用于获取数据。
    GET,
    /// HEAD 方法请求一个与 GET 请求的响应相同的响应，但没有响应体。
    HEAD,
    /// POST 方法用于将实体提交到指定的资源，通常导致在服务器上的状态变化或副作用。
    POST,
    /// PUT 方法用请求有效载荷替换目标资源的所有当前表示。
    PUT,
    /// DELETE 方法删除指定的资源。
    DELETE,
    /// CONNECT 方法建立一个到由目标资源标识的服务器的隧道。
    CONNECT,
    /// OPTIONS 方法用于描述目标资源的通信选项。
    OPTION,
    /// PATCH 方法用于对资源应用部分修改。
    PATCH,
    /// TRACE 方法沿着到目标资源的路径执行一个消息环回测试。
    TRACE,
}

impl FromStr for Method {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        return match s.to_lowercase().as_str() {
            "get" => { Ok(Method::GET) }
            "head" => { Ok(Method::HEAD) }
            "post" => { Ok(Method::POST) }
            "put" => { Ok(Method::PUT) }
            "delete" => { Ok(Method::DELETE) }
            "connect" => { Ok(Method::CONNECT) }
            "option" => { Ok(Method::OPTION) }
            "patch" => { Ok(Method::PATCH) }
            "trace" => { Ok(Method::TRACE) }
            _ => { Err(ParseError { message: "not a valid request method" }) }
        };
    }
}


pub enum Protocol {
    HTTP,
    HTTPS,
}

impl FromStr for Protocol {
    type Err = ParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        return match s.to_lowercase().as_str() {
            "http" => { Ok(Protocol::HTTP) }
            "https" => { Ok(Protocol::HTTPS) }
            _ => { Err(ParseError { message: "parse internet protocol error, must be http or https" }) }
        };
    }
}

pub struct Curl {
    method: Method,
    headers: HashMap<String, String>,
    protocol: Protocol,
    host: String,
    port: u16,
    path: String,
    params: String,
    body: String,
}

impl FromStr for Curl {
    type Err = ParseError;
    // curl 'http://www.xxxxxxx.com:8888/api/v1/terminate' \
    //   -H 'Accept: application/json, text/plain, */*' \
    //   -H 'Referer;' \
    //   -H 'Access-Token: 1662c7f9-08f5-441e-a140-b6646f779dad' \
    //   -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36' \
    //   -H 'Content-Type: application/json' \
    //   --data-raw '{"representTaskId":439,"reason":{"type":1}}' \
    //   --compressed \
    //   --insecure

    // curl 'https://www.ruanyifeng.com/cdn-cgi/challenge-platform/h/b/cv/result/7b18fcdf29ad2939' \
    //   -H 'authority: www.ruanyifeng.com' \
    //   -H 'accept: */*' \
    //   -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
    //   -H 'content-type: application/json' \
    //   -H 'cookie: Hm_lvt_f89e0235da0841927341497d774e7b15=1680437629; _ga=GA1.2.246129242.1680437629; _gid=GA1.2.30939634.1680437629; _gat=1; __cf_bm=nfgHMv2lwocB272t4MEnj7hV_vB6kunHnS8jLKmp.Is-1680437657-0-Abep6seQseoxChvjS5e5uzVDaTjHx0AqW8prA6dQOFCTUYcNuL6/loaxmF2QYNbi2FpBcId3aFP0IxT14uwwGL+LuszqVUTOV+BUF42tds+ECDVJRVEq/M7zNXGlxwHxcQ==; Hm_lpvt_f89e0235da0841927341497d774e7b15=1680437666' \
    //   -H 'origin: https://www.ruanyifeng.com' \
    //   -H 'sec-ch-ua: "Microsoft Edge";v="111", "Not(A:Brand";v="8", "Chromium";v="111"' \
    //   -H 'sec-ch-ua-mobile: ?0' \
    //   -H 'sec-ch-ua-platform: "Windows"' \
    //   -H 'sec-fetch-dest: empty' \
    //   -H 'sec-fetch-mode: cors' \
    //   -H 'sec-fetch-site: same-origin' \
    //   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.62' \
    //   --data-raw '{"m":"yJTUkybzrzpq2VWwEsysOuLBu4mqY3PZzjc2awrk09E-1680437627-0-AWMoBGVFtPBMBHLQC0zjjYtwnwJqCB4TvCLNvmoc+Loq9NtACefR21ulC6G1G+4WN9ACDR0ohw9Z5s5RPqBQluv9JsyUedZACSyEY3t02VnasqNcGgAxiwrrgUWQkfLjYSoehg8neCTCBX8xF1FEXjLaroYXteFmv6vH8n9xfUjS","wp":"HvO7eRwneGORS3DRM636o7buOOZbS7wSuwq6i+O3r6c6RuXOKr6q63yb5iSKzW6B3CXa6E3d7aK16e7Drm36wCS6k8Sei3H8PgxnMtPpGh6AguhWv6UyLmBkypWqmh3Sd589cV0M3O6f6qJj6wf46TsmBR74kRMhO6RfP6Rrdq6SRw76M7RjLZ76nce1pqv6OL76a7Rhj6W365VSi-ekG5c+CJ8zJY8O3ZWgAJvO6eqOl-e+F4WgRDrPADToRRqFSWaz3hm8JO876DMCORzZnK6ZEuwRzZ0Br9HDbU-ztu+HEj+t2P0IRb9s6KDD8u6rAKDhOUjBGGxZD7661vHo3zTCMaIdMYdFNP0eEikdhHuL2BuxzGNu6ZL7PWry3A1SRvuw4bDBefvUS$R8kiSzww943N6w1VrTNO3nxE6KZ8j7rbdDb8nSef7Mo07hp+MPTvu6iOsF-7d6wXARr3Ru3Ko3WS6Do0ty8TISYn01oW4IjPItl+Rg7+-R6bAOI9ic45-M8RAmXREMudi3bR6aYm0vga-eCKzNbOC3H5ug0WcZ6309mbaNs9xBiRcJ-gcrWhtCx+8VfFE1rpN+jNv5JNPDCpArGPd-76l+K3DECRrSjSCgcIVEOrV16qqEbNEAhtuZEllGFL-j34M3iM3dBoTkBl74AZ9evD0xoreMJpSuYtKSx3EPKMPAZ66P2Gma4e1c7W3uNe6GwLjrTuy0RewumhRvcvdm7zeDi1y71ruvmkTOGKuduamhRazejDZq6ZHOZ7bLPP2Lgw30V3FrDr6aC7rK036O6XHSaZ0uWEXSzesHm2D737CeDODFK7nbogvuPC6ORS36ynCRghu7AVVeChc7j7u2jqGuKdu6+pC0g8-bXc7-6o8dDOkejz7RNLknSjGGPr81bPrm0qawKwG337F-svdDz3j2db84AdiG8gm2337x7FujGebOdDGdUKjigCPGzb$gZKg37gPPkemGuhdUkzRRxLzfd2ru3mWHufdHUgTR2tO$0IuH$0OHO+08kzTeiiGhWhmueREB1f4kwD3RoTDOWoOgovCzbT+83kEj8DS32CzaHC4jOidzP33+d3lmWIkH62VmGUDEg7n9dOCC6ZwuhWjiwf2wGIixpjacEwueqw+YI756v+Z5vdZK+d5eui8RePR-Hk93JTup6d0SDP2Z0n2FJ2P04Y2EADMBBe$dZHhoiazryovl5DMaDB$uaK4d-zJi2tHz2RFStyDkg3eknqyHELB12lB23RcGZAydd0uWW$v871Dti4-TVHSd3dCwRwZy2q6kMyvvaAEm$JiARpRcgdYGl1r3druXENsg1ou6RiSP4AjWAY1IOR-wJu31xKwtAfDm74wZD7vyOjUufcB6zRizDylu6tLhqokd506ibocEngGGshz+c62$lGev6hnnRsaHwYZBzP2+Mc2l4hc3RtEn+dLruUtRUcLYIvx7otE9P3d6kgrk+-6$Ho5W$dq-ycPx5Elz3dEjee2TnW5PJ5i3JhrdP9ceOB2x22Xb3B2jaov4XtAejb8gDdOHi6FnVb1+t6FDZS0Awr66OCb+R6t5FzDKar6IW63DGOWOyBdXaRSjCkOhzF27Cdrsuh$uabtGjXerWk48islm3RWKdObgF5FtSPe8bDOaW8vD46dLS0ue0gRgBnbE5+Gfev6whjD8fhDDZa7DcHb6qB8duT0EP60MR84pP0fW7kgAd47Cd6wCR7tSgwdB46Zh0bWxD4OwgZvAC2Aa1MydcP0OAxqdm20ZGLe2r7UHt6g4mev+7Neo4H6u3uxTfyZPDC7AyzvsjS-D1GuPZgvjZD7r7fjwes6SqAwoiBO2i6EtCzPjSvLk3RDHnCLuvLXTD-dpRED7WJ3+XwiM0WfcYE8Kv$O7y6ZMKHKyfC0n2wF+uzu$OSx1bm6GZa$hn7$DSmSOeqjFDeOqO15H-foxG8gAFVG5AmwdsodZnnW0NdTTrH480PCLuVOdKKenAsvs9R+OUdoNaze4DWwT8ZrMje1DYbZma7PeRuz71MDarajjgj$OSK1VCzBbwfHWKvSwI3zwh9omyxmcjhmZhA2WC30ThbZEOOuECPqzwLdTJd1EgqjzT2H2OT7f0Uer1M3ONRy77q1HztudusCj-dP0B71z36$vEWz0Ty7ACCRheYcVBFgCuh$9$S6vwnh34hee8OeWh2fTfSdzJKRt1yKm1zY+U4C6l4TpFSBuXZ0CJk-ETYj+TVBRkd71VHwOvKf4WSwBjiIITwRDMbOHMotu+sESgF6h1rn3Xd8RwdnVHzDCVB4iwiLGbNktvH8DbMJtcoPt88tWSjT9tbE-CntiRETj3rKHHjWR6D+qune4mialVbuSTmc0fabDl4ODwPZgqJuzd0c$5WIosRs1NYTgaTowy4R120ilkGRdC9-zTMmC7m+VCatTyaW1TZHXERa4w3wv3WDu-A3wieBwmcw80bIfy9COjzBaE0gw$OzmnbOi6Svswbqtb5bInkOZEbaGDRWJDz04c4Kx64lCZ5aYo6XWSnKOJFwcuFOWizMW+DDpsZ2R672L2+1R86AKlmV3agjE2G6cBHiOm8S7AmO8t5FgJH6stdKDYfXMp+ucS74VDxD-2pwFiW-ZrPEwztsr1Gk8PgBfMVghGuTCfB1cwr1nuutBBML8RmJDmHRGxm8R4DUGuZxcKRSypWLPAxsq1-dRqzGUBVL1-ino$vUxtKRVehUlsFgtKjhfEJYHuWLhiygBrVGjiR-ogodqGuTMLq32BuwX1dBbSxu4gBc4xCHlRtOmTFw6E4wLTX4SDqB2w-xuzlF7JDsaj+by-REzoy43xdlGh$4x252P8jkm+4W8zDmwzp8FHIZx8+0XdSqHXnzJZxAc8+H0lCPJjO-omGJId46Z8d2S0LJP6xpV6ltuDjnuq1tjHqaj4mdLeJw6YlsL290+m0UEII76gceFuwxGwlF88q6fwyH86wtHW-VXIeotpGfv-YqnwvHw7qpkPXEFWN7BjW6lqXaZLU35sPen8ZzwAr2PO547w2IVNKv7K+lbkc7-LYjJInxxOZ8uLgmOTZmWHCJuyKSVaB0P6Xnoftrx8-0HAlro43Th-$0klWskvVNccz$wnmaVeYg36XlAGbNzHM66pimFBdidvtc0znDMWnEsZqNqMgiPXh6-jOg7WGWSjsv-asKBi$YorOMrjE+ticnTS5vJ9PiJBPoT0NhJzGFnRvEjxNZ4RgbD6wqFMiiU+7M6xj6Y53wRoFdRgcvp6xRoqCqv-8-+KA8yPepXI76OT6Xwou+vpKUqGS+FC13LtdKxa$Js9$pG8SW+1VE5ppldCFdVyhPtpz3Bj7hOg2TUAYJBimr3ZGjH6BN1D5fdmngwr$hj0MYcVwFBNKd8bFYfGT3yye7D7DZ4w-RJ3ni1L2V$9FMuP67vrZ4V$vYfLejR-G73pvYsHu7DsP7m-6$eaKKvamvxhwuDO3SKPWEWZmSKpNgfuxhuS3Dw73pxMhTCvm8DaKYvW6-NvD2RgRGkdD7w7KUv131m3BrD8m1DaDgjaMqeXvo31D$a4MqBR3oKoNVbRvnmjKmWMKpfgwxJuKpvT32eFw6BVPh30MXBJuGeGDoOm3PvFSXeYS$NyKdKpfPfZZj343EqG3ETDkhDyeO3OZhD2wyw4xiD6T0PWeiD1eFwRwh3c2EmED2ehbTxwm865NfA0K03hDUTG3w3AKsTWeqwAmKPhD4kkm22wALwkDIO3m8R063maMgKdwkKP0hKd7Rxz36u6w4fM+PU1RW4a87xPV4ZPG6+martwpFgUZwG6cvrPG0aL17zD87B6SAOySAkSKoC7SARAZPS19DAEG01RuoG0hyEDU14Al4HD4A311PJ04wNFji$DZP8GID4ASTO7G0amFXMhcPbAeweAO73uV8MPqAg0OaJAyvq7IAYTlAZPYAbMvv7bFXLRxAsEpwNFn6zSdGYDIAr1q6c86RIAq8mbq6ab9A1bpAaRG0zu9wXbZwk1z8dZvvAwbbrw9EKUKwWhin6GofzF50hMAF26B6J8pvNbx6-bV6Tm3DGmCw61uXN6bogaEmNT4FvC8CbD5667jTvC9Xg8ErO6vC$XFfhRvCAyPSKmGCCugaHWlDeDZCwCG6ePfMbfy8vC9PS7wC$yUP+MP6vCeCrWidvCHToGHmvCDWhMG8LfRfEuEum6FhDnFyxCqfFDSTli8kS0ZCGXbdz1jE+CmWTkj1gagT7ugLDPn7mk6Pvy2h+iRyVhNChkBAcoZChkrUxCzkCd+igkj1LT+rekn7Bkz874AyNEGCyPcvHFMhYTvfsnoiIFpkB436AyCyZC5kF1die1juuSIXxCialXqXdsq6XSWs9k-vf73FsCv8ToGshM$PC6Usio3PDDhRHWwGN6zsg1iGditsfG6jIv3LHsCWsW6Mm8odkjO6KWfilo7fiaj7Ly$druN6KWmU8uOyoiD31vrRnRy8xscvyTA6j7idj3q6T3G3467dEDuDcvHmNdg6SvH4qDI4pymFbDv3-UBmSUxkP6l3z8kfYhj6lFC3AEN6V3cukFAsRf-3ukysc3lFOsq6VsP6$3gSny8Pdi$3uAq6rPugcungrPZgO6$3duknc3vy4g67h05wGh5o$kZfyGiDoiwgy7HFDskyvg5jHFkg0PowxEsA8nsgxdb7dm40rg0PMFHgfhHW8gh7ygcjsgHWPDqGgUOzCWFmnmhg8ztSLXOjN3nyc35jHWIgfh74sggS1rp60dTzfh+flSLz5j0ztSxw-gfhBrqrq35jZzr4lhpg5osAFzBzjdsShXoXsgJG7yqzmhDw$7xE+dFDWDpyqgtzn4BgJz$GvkGgv+$zY3n+NGJkpz9z$7JkR+fhh+JkNUYfGrT+n4JkS11yz3Hgdgn7ZGV8Tzk+4f3h6zj1mgid4jGafEy8B+36q+Ls34nmH+2+cvldY3Y+pyCsMoq37Hc4CdZvN6$jGap6edhg$jCSwj5+dHy7WH9URH26ld8CWvF13s6kdimHwjG8u+aFHm0HldH6OClSxdgDhRsHpyi4HXp6x3$8IHG8HTWFgUIH4F9Hz6x3p6+kO6pHzu467tUjq6xWeyzgdUfHr+edSkYhedcLjR5Hd87488FDj7DrkH+jL0$8G6n6T6rkcf2isC2P5sU7MkxEUT+yb43hgtN8+t7f11P6U7B7l7swXhN6Jtt7S7ctehGPxhe6bkNh9t5PezO6Y42mVGAD4FG6A7WalTHS8nG6Toxr4r71N8Ma5gO6WMN+JLcHgFcj$dc6D8r+$dny8tp6AteiN6HhB6SlP6$dtSzG46p6p6y4lls8zGqH-lqlfaEzB6yZFHAl9lMm362CidxlzGVCrrzu87XzevRfyIvIRr4ry8TMz6iTID6ULIEI1v3AbARERyll+6FIsSHHuCXlkInUiI5+AIZ7CI4ItIwISUZ7amh7L4o8bIBlHIrtwIZyeyBIsI-oWh3kpIxI$ITdhHLHBmUIqIyeGHpyeIe6aeReLDhDks$8oedemeFssstH74ie7uwIS3iaTeGeXeord+3eDeOejoTote+I8ebesSBeLIwDSdP79i1vplDecekS-6hHmI+6FEvzC6yzoiLIZ2Izd2nAcoAyT6o2RAdij2xUf1C6G2oA36560r37xMBlmn$rDlC8SuA2Lsduru3Ph2z7FmMHBlxGylgUi4f7iald46RhWa6syGhRq2B2C65330Xl22J26FvDsFpyK2tRk8FXMIOB6E6BvBAsdjCGnBX4SB$CjB3FPeoB3XsSA6wse2fBRB1IBCaB5WAERym1gDvunex74tdDllx7hAnynBJDKdmIx7Z26W36stV4h2x7uup2Su$TkSHWHlFDiiXlnmLjIDj1EerxzjhDvovlllWxOxB7DouHvloim6NnKAknh3Dry8UxTOJF3OmFIkP6CxCusxg7g6OI3hHx0d3+2ixx1rVxRfHxEbh7HTGTN6CxJgd8BEgd8E6M6qNnuq-EY6KUZqGaDqYOil7qCm-xNnEusSswSdWIbxGqJeJL$x+xNnTM+67BBmCxAqz6CqG8fq$DTqdisqk7gq4BpyHqhq3um6ogBrI60BO6CxKq87gxYhYqogHTFB5qCxhyOVFqSDRqlzAH6rTxkqNnFm1WCUrV0q$uwbIqkh37ed74ACn76qIDofjq50ub36gaZRUf$zb05F$8yLsVh+6qkV8bn4pZG6jwVbIVqVO7h+NVcV8IxVz12Vn4bdj7y7LyZDDDq72B+x5V+VNGZJDJwhKi9VxLahtq7ixVWyuugL$ylTwBLDED$2g34FBnuoC8s8g8I3pqGV0xZ7aHvVDIJCxVoyCXUe4eVJA6NJ2JwezIor9JKJOHz6oDNJVJLrhC9TVnqwdN18e6oX1Ne8jb5wWNwRP4ybcjhNR8FlpLnUrNDWVnjNevBEZ7RNwNFRENp6gaso6RDl4NG4xdrM9UINAN+EsfP6qNfyDRZ7GnO6cNp69NdR74$NyMm0YhOKXt+6aKsNPNfyY2OH-yoZHD-4ZdKNrKGNp6hLN6$NkLheGK0NmKBLnUyKwK28zbrbdN3u$i6T0I3KwNdjDoIlFN6OUTKjseJLxL3OmUmN3hdNZJu0XbYKRVxKOcV7DomwTNwN9OynGPxK9O9FHS9mNmncYu2GJKq6BL3P8c6OocGciMtKLNdchcjfu7w8jcuuTcYiT79+NdpvRygcKdWbsS5lvyecL1iatcV6gc7myTz1F2Tc3O31zPEc9O-K1v466-mUga8-1v6Or-P7V0lKicR6iUnmoyIDFoz6TcikwcpKPcuGofO-M6+6n7edTc34z68AcUTRTcJtF0P-g2KcKbAAj-9O8G48LySvcca+v07cPc8Ghn9-Tcx8ibV-R6n6ovF3O7wIncvn4cY-icb1RDktBPEcn-G9KUND+HG9zDSdfaxKCh8AwExKi-xK7KpNdNs9B6bKP3X-idg9N8SvChU9iaxKh6mjJ9Sv2M4-EcZycIc9gDNmEKwNEzaJkPEcSKUta56O85dK61y5BKEcHh-vsK0c6O2-I9wNm5bKkitSe6z4j5q7JRN8xKoKw8WKfMTK+gYNwNlDfhu7x9$4xK4KFKhLxKdNN5WVN5I-R9yhNPEcXKbKRTN5vvgTDuN5qlOD-Aj-n6zAh5U-6OobEdy1V9Ecjpq-G5fpW5R94pEM0lIcGpgTS9LpKKnUy5PKdUbK2TmcnZ3nFKzuDcR9O4vbTNIDWFNJ$8hyrTmy1JYNGy5p5IYpC7S7EKptC6$wyBfySYaYvYm7yBoiEAl7dlmBgUEA4xP+AEMlr+LRovDfkikStSafq7$+diPDlwBmCJcp1l4-lqLN0ndxjoS0e+EKio04Iy99qYMoJYgybq5flRcYXoNY+66C713YmiYNOvp6F2bMhnEKUnfyu$cpd+ndvKM$p8xLaKoZ48B65eVyoyNXtYpK4$DuX$rTF$RfN-E$C$6rKdzuo1lIEKK4R-AVy8Youh37blYf8nbf1hKNFfuzpKY$cOevk$t6U2I5$LR7iPgAJuhgw336Hwu4klr6LP$gR00B2oDdyT43rwAmJNZkMLNDktW4ce7kOZEKj-Wv3OmorsLNN5IVtg074nnfKITMe0dwNvz8m9In6Y6DzP4PJ4ppP7lH6OWfsJu2xKTOa577r6PON6l-CObPq$b4SP6l-VOiCAitswTbT3cfczshRPEPX3ulg6-BhI1h1MUKiVE4SfRHiaRJeHbmrKLR8v74Talk31DXue4lpU6316WRje-wdATKmD6WRERt4cjR9c0LyKawBUdUzG8Ly2Zoe+wwlYCpviesRr06KYUuzuwaxVuaR87sOk+mEWTty0TuVOB1PZwz6r6OaKOk8bUVsoTnbuvILdbn63oHjSCUuJJa0x1w2vRfuSl8RG3DwR0AuxSddiPv9kL0PuSWTi6Ww23$6oiTSb8DXeHeb0cF3KcAmXAJ+FDyaWdXwR+Srwsa9URZj1rv6-s408aL$r1iGWo67ZZ3BSkriiWGlyZ6rjH11O4K016o4aNOm06Wei6SanmRvPAESeIPhZv$PvnoTZRpRt7h7evJsLKN-Redta$vZRq7AMzq09aAsegeFDFO6PvDODJRrF3-ez6Ha1vknKPm72rPEO86HNwEexgHMAOg7RX-l2HRZ7msznCybC3PbwSoZbNKjTbKWuG1mMJpwO2i64eE64W57wi21koWZ4mdBpRiMmBjPGjzdi1cfWgmGe0P5Ei3CTb3mz1qO+8PBevE335OiGBmZlm-we0RwKM8HO7KhO$0Ou2Pp6rY3tMD3rgNmO6Zl3WTu2ioR4O6OPSo8ZV169P36PW0pLmh-dW6u26Mqs+qWPaX66RbWtCueK02$Sj7ByUmrOAAW7O6ODZ51DqbwyjMWZsrPMrqk4zrAOa3Ypsee3bWP7bkwlD4iB3T2oTydTrn4b4O4OoACz5r4LTSwz2crBZait0wlBiynSsZ8Zb8RkmFIUApcTshfOSe320A-a7-Ric624DoOUZc0V9wVKb2+Uvednn4ORKDA+S+Jq1VErkuvSfT9U-1JfBq7Mab8T9Ko84TrjdC88Jm5dCUrCPO6aGvD-tS4qKmbuY05Z1n-o9zkV2pAknerh33BAhu8OM79olTxq2UsXO7ydSTDR2uMAubPZpoYVwxCn6ywuPTKqSMcvmH-Pk6qRJG+dRC+h6W4ByJ05yJvBRbz6Pq3GOywE00wcScpS4rX2-P5PXXDHp+PZLdSEbWJz4rLM2GduZEYhK-bAHlJIza0whOjZG06eMw3RM0KTHbOoTwPb6CBijU$0SRTBgbw1HKTwcZeMj+bcqktc5j2P0vlbkAUpLJ2IBM6LiRheiKtjOR4j5XXJGqePc18ea1PDOXMl7e7CJss6c8S6Wi9KvFBVElrlGemT3uKyVW6UKMPdwuwz6jSU1UbRl3pCrZSC6MTWeCOaeDZTbW4OdPsqW8RwCGkKlTdh5jLdyHirrUmOjboLgBvj3Dkh0sbDC+jCb3aGnXnd9OFjmCd+wli+bGgBdgdtSVhddroBFuYbu9mhHX9jcllPZL0lTz5L87FjfOhRUGZT8-eL7OCny26T3b4gjSbP6ys1BgTG8Zwm6Y66OUgqgGzvYbkySX0G7X8j2BCMXlTPxjXxBghCsdoD-bqEZ3ycCrP3fJzTLVCBcU$ycRiiRFmgC6kBM3yBsHNDkgRZ7Yjm4iiB3sjDCEIkTP73NDwOqb$WuE2HDxNFuVZvUbVC+e$B16N8LgO7wI320DwFdovl7DwD07Fvu6hr6F8O6pk+PWCP6PLjAmMwhv9vV7JnUa3P0USYhoim2hi64K4BFk3oT76t7NMY7bkSrILG15salCJ6oa+rcLjxq36kPJux4kIgH7-w4ig6R+T3Dw0vPAUdiZz3bm+gD6qkPSHSuJrfuTtASy8Dy3r9W7oAw6P6q4SBKZRcusaoideKW7a2JPJWy8deOcxtbAIwRN6riY1jZaluavn8nqw+Kfu6uW6R8dLeycZe47eOgVqiPZ6rW4wE6FFRi1-izTTCCOJTluO3yt7sPTsvsR+De6S09ze447B1J-TfkfAE6UPOhg9WDRUsvnUPbi380rRcb5fN6dawY6aqgCvnC7JiBIdyAlP$OC7k$eczrRBy4kpyHLj6e-gX8XkNAq4KxTOeHCwEXnSFBY9JG3xK6Oo+awvw2GcMu6B7mHCYOY8G6AkdxluwCAn290e6uInnRcL8bE7+mWu0FJ$HtT8wLnuRp+TS3$cIWeODR1yr6p7M4kwjtBBFH-ut2qZW6EaPEPtWOe6M7OffeZKY4kgK1+O4o$wZuhfCwRT3gTwgV9AE3rZG3tD6MPP69g8TNDN6b0nCrgAGsi3Zva0lu+rNU2Us5R8b3yqvdDqTdg$ROTGcoRww$OBRfPda38-Mq61WWv$c5mk2WeRfDWa0dShhVGq41illP4Z3sWgPe3ltKTDapDzD6$62r9x7Ug14F6Pkz6q6TWgAWo64s57-aTwgA2MAdDY7deK-Wt2as6gVo2igVwRLXuy1JWMuTDCodG6KU+6nMXaeoe+3dG8wPR5VWqAF6UO61jnGcwouX6ihYh-bjatDC1zG1MzbABpb00AdDhzSbTpDtnk-16DflzbJv6$LAT53-WpaH6hc6lqt7Ujxh3u33WuALZu7DCYseoRL77OHieCdapJa3eAOTjWrK2OH7qMWkGMhYcVD27gwASoZJE41tPPC3nvWaOpRd63cGreZoijszi3S7pIb1JDeKprDgLpTi8UMsAPJ3zT-qmydvbuhtc46qvLrpMtALu3X-iTORevk7-0ua2x6Sz36m3Gi7a2swr7oT7w6FKuj+7SjHmjsuWVGua8ozmTW76RtmMoO8ocjOgnX+Abdhgfowz3UWhPbusGz$DPMkdyVa1Vdz7WcgCMh82T+JbFO34uF6xvtSA6+BkOdDAkziemzevl-Coh3-syx6ouvwjdmH8Mw9kRfDSKz7LAlr-gDZJzU1e+jfVAridqzlkbJxjutzuuqj6jt3Ae6kZgTq2hvKlS0IIzuujg3SadGGHgmgh9a6fCZ6ZoSi5Wq1Wd36A1q7Zuadeb7dtknxO+55TtTSDtGrWsSyH-+zmp7Tbh81vqRriLCiTkV7re6zH-mS+vNwo+17zHuMCxdG6kFq6BsBK5qfvVY6Ov5hUnSG7uVL+9R5EdrmJd$os6BjSZ5bePjYqaz5yOEmV3IdluAbtG7kRJOeag6iSR6ZiD0mK4hwFAGaBsd3M1WaBd+1fp4l8RmRJVZbi4q4chPTwP7kGs8SZr-TmwHCjoGPOZH9TDmzCPnxdn+0ntoR5VJz0rKd0S3DZ5h8vtDsIZBKE8gw5+umNGqwqPPPoOaZmXPGqWPj7G6Hoz4izp8u9dw8HsYBfeYO26BM6Ddoa+dr3ccsxCaB6nsZP1K7muHgFjVkYWm-gLur4CmCZiPhnv8JxgWvn2k5zN4OhhgDTx4m-mlvf2m+BRo7WARldPSaFuab0PP8l4Ju06qhWiWdKKRfsYzjLMgTsetPNw6vNvmjTXBEwcudai3lC8ZK0AdO+M2i663Tk6E7igeEJ4MdegIDA-yluwSB8x1xbda1gDYR9kgXucCMrL2icBu2MsLZPf5lRJKscDehUY19eb7zX4MBE6TbJKOx7$bV64RM4saYia1R4u06n6d2YegMEGgE6XDL4it6HqR6OEn3GImvWOwxvOyu6fjK5WZaHG+Og8YMTD0mG20dgreOLC-8u0OrtoJTBCT-w6kjwz8vDbRR+C1eUoTD$8UBJYHsL20vnf7Ouqmrne1hbml4aawSXpNNIe6ajiewH1n60xzgO6qiDO6iycGAZP6mChdHtNasBu2SSurfb$mWoImEDCBunXS66wGBkcVDFEf7CuZ6bOqhDewr0zadSYZRN3+wufigzrIUqkV1zkSNMijsAnMASAEqcybtUb-OKKxls0aY6eCZl6f45OqOji2C-9W12ATjmj1kKBedrV2PpmWFka-v3Cy4dAMgiEhLtSf-NkchtE2CG-w8khPvgdGrAXRVsEsrLrYj$9CtDesITxbKj$nt9aUWgqigmECAYt1m-1W1DRhR-JvmFhXTs6iau42rpsAL2ePCWGo8v1ceX27+RM$0PfVITDxrThXNAdap0qtOhOlSq2u1tOc20MbToZWdvhReFPIg4XXN8rw9Ce+EdvU+3Hl60d7PU6hRjsZAt0CjuOzRabZiAS2wm0aE8J6mzct881DtzZK-7wRpW78sB6suA5JvR+V1SH7MUSZe-SSWZDxgBmWrTl3qdHfvoCUWgZJh8T0NX+506lGLu04eVvR9UDSmkep2+fOa8F8t--VOyxNeSDTCVPMtoeRN2mpmAzCp177moGzTX+6JR70-vP31N8iDwiBO6BS27sFAinXK$ej66OTZ3i-RKekMGDaTZ63T4WqMEwZM8jfCuxH9aDB1OA2bvn-E5k3M7ek6q66JSE6kaPgtXSRw9dqJwNhI1LEJPm3O9t+-zFO4IZXRqOne3HWh5crZKBBuJouBRMLC7ZdEF6L3Ba1BUKTSHww6WrfuLD06fjb0Y6qOYrmijdJ4h06z5iDTDRi6L0i8adVEzKxE6+r9I16ehecEIBlWWsT7UBu4bgySfutVfNbPyp6W9nFlp6burN95bphmIhw15XSGE7eVpyeRy7MdRdBhHfsch0Zg8VrrP7L3ZLFBGZYDHD6y8TZrJjw3b876AtG3G7DEgiKFKDOCRt28-2m90KDss3m23DZOjRNZZY01JGWcw4aaYo8PARTDDSx8vXMvDBuXiVO6TdW6FPvETZejHGaLy2GKTTdYK-CxJHtP8ErDZtjHKB6EabL8-HO63mBZuUCJU18fHdx$OvDeLYfHRirxDjhRZZ$P6LrRwiZx6BSADwYvqKn2MAF44rn6S1wNtSfunVG62GuAHZevPZMGPTrtul7TbJsBdliHku6WS7uuMnYBV-KLyFfd7orGB0Kmp3vD-75M7wuOqi2PS-Ur9eWHiClvHSYVH98FH2guB$25gO3V0dcmif175v+Hcp-ac8BkASZvBRLa-AH8hbEurPSArOLm6YmMvA9dXE0mHSuTZZNGrfrOgMvLPaipjkxhVR9ncPc3Op-Cc7PoUogD+luRX4V2aOxGC7ioU827uW2118vEum1cvfB2AiVRj-OMMLbJ+YFoRrBt7LlmxOZ7DiKOm35hc2bbACLXtKkifLVZfylWmeas5LaaoAb9dAPwiKSLFB2WyOE1mXhlKcB7ziuHuRxZcKVu9+idffGw3EwbbvDKy9yUDJFcphnoorFSm$cUL6xFbhrxU81xwL6xvvHRkyZ16hw8MfkoETYCi5G6T6mw8dIDjz2uqImcPLuHI-qrwNu2DAa5XwetHY7P22iU1xMJKefHsBek1TdmUWTtc1BPkx1+KLO9gc4xrkL1MiksZtLVm-8nnqwiAJeLCxGCpr+Z6i5IcLLclDvxB0+rJcULAs6HmV8MiiHq2Ao1lWnnYcZJu2f7ZZuyY-Cx0wwurncqG3HR7m7Bf7R7Oo8Mipc+yK6Am6st5iAM+O4sWoNKiticL2W73VtbDS$5iN7Zi6MgbyGA3VO2gSPv$ZLJcrcKOZ-DU-eaLWD6-W3cPa4wjgF7Sc8qi3oRr85grD-CoDjsOodzOwNLtSC75C8RYGHfcu2ASlRm1K6X8wbaWFJoIAJt9lSD2DPxSyaVls8n+HgbFdfbUWr7-u2BRMWv-4Op7Pdxk7k1vnXz+BoHGb7$D3VKLex7nGGYBXYcCSWEo2exh8axaB5dSgvOSxY6smR87OJY650cojdKWH8lWqVg1OybkWij4V+-g12Vbno1P1n2pAWrTylTWj6hIvzXnTGhntlmM2I6ujL3YWt+KBktLxMIyU-LTCa1mz0TNyl2gF-w2YgC1ozUL5zc5qvAAqqrGEexK8V73WhM1JhpYqYzaA2Z+1UrSBSI2EMZbHpM4H+-XA3AvhIEqG+e4x2EdbvMJKY4t-OVAztvIjJhpWcoucANvv+G4XZcDBMza8AsJWtJ-yBYga89EYEtBkvxL+hJO3kEX6kIVma9f5zJWXyLnOd+aEZI2oXUcdHnz9E-w2ot1LohMzoAmZw5HV-7TsgL5Vll4KmmZFV+s6sMlB7eLKBqzXjDgJ4XjL8OyCtGNTYJHF-WljCW3jhlrjqL+VKPsTaGYJS$4RJZUdy9KMHHDEVBjPETFw6Xia99vhvxMMmY47Z98VVHpq3elOR19eVKH7rsmYspYzPxqHXrxMVWX19oJStNfWClmiU4k6cd5MA4kFzA4b2jtAqOK9gXtLjVMz7n9fl+XtLXJjt91a60AHV4wxsiymxEGc58NOJAthnvqYH12c3V4thnMlGltBLhOrs0ULn7H1qcqrotyaq7GeC4Eh2qBKmfuD95x4+VDww5jXGtXKL$qAsFmoEgmPF4xx2iBMsUUj2C4vVis-5hRcpr3ocxwJzYZ2eYH4moNk7qo90x4tImYM4l$o+jJft+97qtHIiEwU2tk9y2B0XXNwJCtZavr0MAfM0lciFE+Vm3FlpeZd194oqMzLUECmEpZ41d0lnW9sZ-1u4$ecgIhx1Gtp14rvc3VX93kv1In$qygNMypYvXjqwqvgVoi5X51R4AqldnfoCZDt6EsHZYTEyapYpMGNJMzFjfPGUtcEyxZzM8tu3b4Dh+Idvz-cpSmyNw5gH0NPSSAI43f+IVkAWh7Ipz5LvrHgCJq-kTeiwdvhefvMV+WpRJIqfai7AkJ9yXKaJl7zwvaGRO5kHJp7P0oU5hCa046ogzub666eypeonOOIoebefOtSzRC60P+ePpeDw2BdueF+KwsHDRHR0ve1eX6I734d9OcjT0e0Y6wMxZDGRqzNwAGAw2JeSvi4miW6BO-am6666D66R8acfHz6H2aSeZfBe4i669eZIyZXapKuu3E66DSJK1j8u5UbZRo5uldS-aRRR3eEi2DZw67eZNnW5OJKNUGLwZRRUsaXL36-0NLRje4iRct2$1UmEa7PiuudBLbxZvuRUqncXZ6C$ij3esi0uGvJAunpZYZOI7CXb7bISh65NZ6Y7edF1Hgjl5XcpB77TOKmU8w5t6OjvbroFyBReup7uk+hZV0NwlmhvPkvWPZ6vSvuGBhM5CpflVfj8tBiDwSfynP5VKI2AAM$wnla69KZ6BAIvNORFiJK$2FDDDsRnfrAAzvobsDFWUAiNKoBL6S8Fjndwdc48Yks8ooysioc0BMsodIKeD6Mc3BEbIdvZR+5ocex0CyysGlF6IldwaMK16w6CaCNH2vfDZ3oRqiJZd$GRK1esB90O6$TUWR6WV6Eqoud$u3B1RtLyJgCulXDDWeosfA+Sm6t6JDDvUTjUK$eaPFimDn0Aj6Z2FTmOD8Uo5Fd37l3$1DwEw0UKZ6qzNUGDRYJB6fJ6F8ESRqOuhaTC5l-RwX6d0sKzZPZ$cwZB2o8IREC+Nve3iv+51bI5ZixeAPxWN-YMyjpt68rSgUKV6uO8W6U5xkBqUaiO5fsR+a7rx1Zh1DGKL2WRT1Gd5ZHTIgUW5AKjVsL3O$-RF6CF7uJbK1e6ea0t-Fa-0Ah$E677JKmVM6D8K0RuwI7cwatN5pBlKr5eO666runXMzjOTr6Z$wkN6Cx-CBk3432pAdUVnOlhsaWV6NEYmOmf7abZlKJmg9rNFO663o766","fp":{"results":"cb5fa63e0b921e6276ac44abe03cb73b"},"s":"0.47154076090002767:1680434905:r_38p-hmFyg7qWNM-9HhSfeIKEmXRB4qz2k-Sli4L-8","t":247.40000000037253,"src":"worker"}' \
    //   --compressed

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        return Err(ParseError { message: "parse curl string error" });
    }
}
