import { useState, useEffect } from 'react'
import { message, notification } from 'antd'

const Screen3 = ({ current, wrapSetter }) => {
    const [state, setState] = useState('pending')

    useEffect(() => {
        if (current === 3) {
            setState('running')
            autoInvest()
        }
        else if (current === 0) {
            setState('pending')
        }
    }, [current])

    const openNotification = (title, mess) => {
        notification.open({
            message: title,
            description: mess
        });
    };


    const autoInvest = () => {
        setTimeout(() => {
            openNotification('Auto Invest', 'Success')
            wrapSetter(4)
        }, 3000)
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
            Auto Invest
        </div>
    )
}

export default Screen3