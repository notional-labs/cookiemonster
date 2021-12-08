import { Modal, Button, InputNumber, message } from "antd"
import { ArrowRightOutlined } from '@ant-design/icons';
import { transaction } from "../helpers/transaction"
import { getKeplr, getCosmosClient } from "../helpers/getKeplr";

const style = {
    transfer: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between',
        marginBottom: '2rem'
    },
    transferInfo: {
        padding: '10px',
        border: `2px solid #c4c4c4`,
        borderRadius: '10px',
        width: '12rem'
    },
    container: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
    },
    form: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
    },
    button: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        marginTop: '2rem',
        marginBottom: '2rem'
    }
}


const DepositModal = ({ account, wrapSetter, cookieMonster }) => {

    const success = () => {
        message.success('Deposit success', 1);
    };

    const error = () => {
        message.error('Deposit failed', 1);
    };

    const handleClick = async () => {
        const { accounts, offlineSigner } = await getKeplr();
        const cosmJS = getCosmosClient(accounts, offlineSigner);
        if (cosmJS != null) {
            transaction(cookieMonster).then(data => {
                console.log(data)
                wrapSetter(false)
                success()
            }).catch(() => {
                error()
                wrapSetter(false)
            })
        }
    }

    return (
        <div>
            <div style={style.transfer}>
                <div style={style.transferInfo}>
                    <p>From</p>
                    <p>{account.address}</p>
                </div>
                <ArrowRightOutlined style={{ fontSize: '2rem', marginTop: '15px' }} />
                <div style={style.transferInfo}>
                    <p>To</p>
                    <p>{cookieMonster}</p>
                </div>
            </div>
            <div style={style.form}>
                <div style={{ marginBottom: '1rem' }}>Amount To Deposit</div>
                <InputNumber style={{
                    width: '100%',
                    height: '60px',
                    borderRadius: '10px',
                    border: `2px solid #c4c4c4`,
                    fontSize: '2rem'
                }} min={0} size='large' />
            </div>
            <div style={style.button}>
                <button onClick={handleClick} style={{ borderRadius: '10px', height: '4rem', fontSize: '1.5rem', backgroundColor: '#9b8da6', color: '#ffffff' }}>
                    Deposit
                </button>
            </div>
        </div>
    )
}

export default DepositModal