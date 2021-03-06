import { message } from 'antd';
import { WalletOutlined } from '@ant-design/icons'
import { getKeplr, getCosmosClient, } from '../helpers/getKeplr';
import { transaction } from '../helpers/transaction';
import { deposit } from '../helpers/API/api';

const style = {
    div: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center space-between',
        alignContent: 'center',
        backgroundColor: '#ffc369',
        borderStyle: 'solid',
        borderWidth: '20px',
        borderColor: '#ffb459',
        height: '35em',
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

const Register = ({account, wrapSetCookieMonster}) => {
    const handleEnter = (e) => {
        e.target.style.transform = 'scale(1.01)'
    }

    const handleLeave = (e) => {
        e.target.style.transform = 'scale(1)'
    }

    const success = () => {
        message.success('Deposit success', 1);
    }

    const error = () => {
        message.error('Deposit failed', 1);
    }

    const handleClickRegister = async () => {
        const { accounts, offlineSigner } = await getKeplr();
        const cosmJS = getCosmosClient(accounts, offlineSigner);
        if (cosmJS != null) {
            transaction(cosmJS, undefined, undefined).then(data => {
                deposit(account.address, data.txHash).then(res => {
                    wrapSetCookieMonster(res.data.Address)
                    success()
                }).catch(() => {
                    error()
                })
            }).catch(err => {
                error()
            })

        }
    }
    return (
        <div claasName="container-fluid" style={style.div}>
            <div style={style.buttonDiv}>
                <button
                    onClick={async () => await handleClickRegister()}
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
        </div>
    )
}


export default Register;
