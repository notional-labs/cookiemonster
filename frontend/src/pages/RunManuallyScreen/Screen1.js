import { useState, useEffect } from 'react'
import { pullReward } from '../../helpers/API/api'
import { getOsmo } from '../../helpers/getBalance'
import { message } from 'antd'
import { Tooltip, OverlayTrigger } from 'react-bootstrap'

const Screen1 = ({ current, wrapSetter, cookieMonster, account }) => {
    const [state, setState] = useState('pending')
    const [show, setShow] = useState(false)

    useEffect(() => {
        if (current === 1) {
            setState('running')
            fetchReward()
        }
        else if (current === 0) {
            setState('pending')
            setShow(false)
        }
    }, [current])

    const success = () => {
        message.success('Pulling reward success', 1);
    };

    const error = () => {
        message.error('Pulling reward failed', 1);
    };

    const fetchReward = () => {
        pullReward(account.address).then(res => {
            success()
            setShow(true)
            wrapSetter(2)
        }).catch(() => {
            error()
            wrapSetter(0)
        })
    }

    const getAmount = async () => {
        const amount = await getOsmo(cookieMonster)
        return (parseInt(amount) / 1000000).toString()
    }

    return (
        <>
            <OverlayTrigger
                show={show}
                key='bottom'
                placement='bottom'
                overlay={
                    <Tooltip>
                        <strong>{async () => {
                            await getAmount()
                        }}</strong>
                    </Tooltip>
                }
            >
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
                    Pulling Reward
                </div>
            </OverlayTrigger>
        </>
    )
}

export default Screen1