import { useState, useEffect } from 'react'

const Screen4 = ({ current, wrapSetter }) => {
    const [state, setState] = useState('pending')

    useEffect(() => {
        if (current === 4) {
            setState('running')
            getBalances()
        }
        else if (current === 0) {
            setState('pending')
        }
    }, [current])

    const getBalances = () => {
        setTimeout(() => {
            wrapSetter(0)
        }, 3000)
    }

    return (
        <div style={{
            height: '10rem',
            width: '20%',
            borderRadius: '10px',
            borderWidth: '30px',
            backgroundColor: state === 'running' ? '#8abf80' : '#e6e6e6',
            lineHeight:'10rem',
            fontSize: '2rem',
            boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)',
        }}>
            New Balances
        </div>
    )
}

export default Screen4