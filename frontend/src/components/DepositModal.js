import { InputNumber, message } from "antd"
import { ArrowRightOutlined } from '@ant-design/icons';
import { transaction } from "../helpers/transaction"
import { getKeplr, getCosmosClient } from "../helpers/getKeplr";
import { deposit } from "../helpers/API/api";
import { useState } from 'react'

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

const defaultAddress = 'osmo1vxgcyq7nc8d8gykhwf35e4z0l04xhn4fq456uj'

const DepositModal = ({ account, wrapSetter, cookieMonster, wrapSetAccount, wrapSetCookieMonster }) => {
    const [value, setValue] = useState('')

    const success = () => {
        message.success('Deposit success', 1);
    };

    const error = () => {
        message.error('Deposit failed', 1);
    };

    const handleChange = (value) => {
        setValue(value)
    }

    const checkDisable = () => {
        if (value === 0){
            return true
        }
        return false
    }

    const handleClick = async () => {
        const { accounts, offlineSigner } = await getKeplr();
        const cosmJS = getCosmosClient(accounts, offlineSigner);
        if (cosmJS != null) {
            const amount = value*1000000
            const recipient = cookieMonster !== '' ? cookieMonster : defaultAddress
            transaction(cosmJS, amount ,recipient).then(data => {
                deposit(account.address, data.transactionHash).then(res => {
                    success()
                    wrapSetAccount(account.amount - amount)
                    wrapSetCookieMonster(res.data.Address)
                    wrapSetter(false)
                    // localStorage.setItem('COOKIEMONSTER', res.data.Address)
                }).catch(() => {
                    wrapSetter(false)
                })
            }).catch((e) => {
                console.log(e);
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
                    <p>{account.address.substring(0,17) + '...'}</p>
                </div>
                <ArrowRightOutlined style={{ fontSize: '2rem', marginTop: '15px' }} />
                <div style={style.transferInfo}>
                    <p>To</p>
                    <p>{cookieMonster !== '' ? cookieMonster.substring(0,17) + '...' : defaultAddress.substring(0,17) + '...'}</p>
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
                }} min={0} max={account.amount} size='large' step={0.01} onChange={handleChange}/>
            </div>
            <div style={style.button}>
                <button disabled={checkDisable()} onClick={handleClick} style={{ borderRadius: '10px', height: '4rem', fontSize: '1.5rem', backgroundColor: '#9b8da6', color: '#ffffff' }}>
                    Deposit
                </button>
            </div>
        </div>
    )
}

export default DepositModal