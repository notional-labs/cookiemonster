import { useState, useEffect } from 'react'
import { pullReward } from '../../helpers/API/api'
import { message } from 'antd'
import { Tooltip, OverlayTrigger } from 'react-bootstrap'
import { checkReward } from '../../helpers/checkReward'

const Screen1 = ({ current, wrapSetter, cookieMonster, account }) => {
    const [state, setState] = useState('pending')
    const [show, setShow] = useState(false)
    const [value, setValue] = useState(0)

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
            getReward().then(res => {
                setValue(res)
                setShow(true)
                wrapSetter(2)
            })
        }).catch(() => {
            error()
            wrapSetter(0)
        })
    }

    const getReward = async () => {
        const res = await checkReward(cookieMonster)
        if (res.data.total.length > 0) {
            const osmo = res.data.total.filter(x => x.denom === 'uosmo')
            if (osmo.length > 0) {
                return osmo[0].amount
            }
            return 0
        }
        return 0
    }

    return (
        <>
            <OverlayTrigger
                show={show}
                key='bottom'
                placement='bottom'
                overlay={
                    <Tooltip>
                        <strong>{value}</strong>
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