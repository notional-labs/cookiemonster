import { Typography } from 'antd';
import { useState } from 'react'
import { WalletOutlined } from '@ant-design/icons'
import {
    SigningCosmosClient,
    coins,
} from "@cosmjs/launchpad";

const CHAIN_ID = "osmosis-1";
const URL = "https://lcd-osmosis.keplr.app"

const { Title, Paragraph } = Typography;

const style = {
    div: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center space-between',
        alignContent: 'center',
        marginTop: '2em',
        marginBottom: '2em',
        backgroundColor: '#ffc369',
        borderStyle: 'solid',
        borderWidth: '20px',
        borderColor: '#ffb459',
        height: '54.5em',
        width: '30%',
        borderRadius: '10px',
        marginLeft: '50em',
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)',
    },
    buttonDiv: {
        alignSelf: 'stretch',
        marginTop: '5em',
    },
    button: {
        borderWidth: '0px',
        borderRadius: '10px',
        size: '10em',
        backgroundColor: '#ff9e61',
        color: '#ffffff',
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)',
        width: "80%",
        height: "10%",
        padding: '4em',
        paddingTop: '2em'
    },
    p: {
        marginLeft: '1em'
    },
    content: {
        justifyContent: 'center',
        alignItems: 'center',
        marginBotom: '10px',
        fontSize: '30px',
    },
    addrDiv: {
        marginTop: '10em',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'left',
        alignContent: 'left',
        marginBottom: '10em',
    },
    addrContent: {
        backgroundColor: '#ffffff',
        alignContent: 'left',
        margin: '5em',
        marginTop: '1px',
        alignItems: 'left',
        borderRadius: '10px',
        overflowWrap: 'break-word',
        padding: '1em'
    }

}

const transaction = (address, cosmJS) => {
    (async () => {
        // define memo (not required)
        const memo = "Deposit";
        // sign and broadcast Tx
        // THIS IS A TEST ADDRESS 
        const recipient = "osmo1cptdzpwjc5zh6nm00dvetlg24rv9j3tjh7wnnz";
        const ret = await cosmJS.sendTokens(recipient, coins(1000000, "uosmo"), memo);
        console.log(ret)
    })();
}

const Main = () => {
    const [loading, setLoading] = useState(false);
    const [address, setAddress] = useState("hello")

    const handleEnter = (e) => {
        e.target.style.transform = 'scale(1.01)'
    }

    const handleLeave = (e) => {
        e.target.style.transform = 'scale(1)'
    }
    const handleClick = async () => {
        let val = !loading
        setLoading(val)
        if (!window.getOfflineSigner || !window.keplr) {
            alert("Keplr Wallet not detected, please install extension");
        } else {
            await window.keplr.enable(CHAIN_ID);
            const offlineSigner = window.getOfflineSigner(CHAIN_ID);
            const accounts = await offlineSigner.getAccounts();

            const cosmJS = new SigningCosmosClient(
                URL,
                accounts[0].address,
                offlineSigner,
            );

            transaction(accounts[0].address, cosmJS)
        }
    }

    return (
        <div style={style.div}>
            <div style={style.buttonDiv}>
                <button
                    onClick={async () => await handleClick()}
                    size='large'
                    style={style.button}
                    onMouseEnter={handleEnter}
                    onMouseLeave={handleLeave}
                >
                    <div style={style.content}>
                        <WalletOutlined />
                        <span style={style.p}>
                            Register
                        </span>
                    </div>
                </button>
            </div>
            {address !== "" && (
                <div style={style.addrDiv}>
                    <Title level={3}>Generated wallet address</Title>
                    <div style={style.addrContent}>
                        <Paragraph copyable={{ text: address }}>
                            {address.length > 100 ? `${address.substring(0, 100)}... ` : `${address} `}
                        </Paragraph>
                    </div>
                </div>
            )}
            <div>

            </div>
        </div>
    )
}


export default Main;
