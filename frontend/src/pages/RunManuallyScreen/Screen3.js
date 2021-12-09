import { useState, useEffect } from 'react'
import { message, notification } from 'antd'
import { autoInvest } from '../../helpers/API/api'

const Screen3 = ({ current, wrapSetter, account }) => {
    const [state, setState] = useState('pending')

    useEffect(() => {
        if (current === 3) {
            setState('running')
            autoInvestProcess()
        }
        else if (current === 0) {
            setState('pending')
            wrapSetter(0)
        }
    }, [current])

    const openNotification = (title, mess) => {
        notification.open({
            message: title,
            description: mess
        });
    };

    const error = () => {
        message.error('Auto invest failed, maybe you are trying to auto invest more than once each deposit', 3);
    };

    const autoInvestProcess = () => {
        autoInvest(account.address).then( res => {
            openNotification('Notification', 'Successfully re-invested OSMO')
            wrapSetter(4)
        }).catch(() => {
            error()
            wrapSetter(0)
        })
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
            Auto Invest
        </div>
    )
}

export default Screen3