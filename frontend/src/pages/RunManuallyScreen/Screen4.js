import { useState, useEffect } from 'react'
import { getBalance, } from '../../helpers/getBalance'
import { getTotal } from '../../helpers/convertPrice'

const Screen4 = ({ current, wrapSetter, cookieMonster }) => {
    const [state, setState] = useState('pending')
    const [asset, setAsset] = useState(0)
    const [done, setDone] = useState(false)

    useEffect(() => {
        if (current === 4) {
            setState('running')
            getBalances()
        }
        else if (current === 0) {
            setState('pending')
            setDone(false)
        }
    }, [current])

    const getBalances = () => {
        getBalance(undefined, cookieMonster).then(balances => {
            if (balances.length > 0) {
                getTotal(balances).then(total => setAsset(total))
            }
            else{
                setAsset(0)
            }
            setDone(true)
        })
        wrapSetter(0)
    }

    return (
        <div style={{
            height: '10rem',
            width: '20%',
            borderRadius: '10px',
            borderWidth: '30px',
            backgroundColor: state === 'running' ? '#8abf80' : '#e6e6e6',
            paddingTop: '4.2rem',
            boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)',
        }}>
            { done ? `${asset} USD` : 'New Balances'}
        </div>
    )
}

export default Screen4