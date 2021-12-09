import { useState, useEffect } from 'react'
import { notification } from 'antd'
import { Tooltip, OverlayTrigger } from 'react-bootstrap'

const Screen2 = ({ current, wrapSetter }) => {
    const [state, setState] = useState('pending')
    const [show, setShow] = useState(false)

    useEffect(() => {
        if (current === 2) {
            setState('running')
            identifiedBestReturns()
        }
        else if (current === 0) {
            setState('pending')
            setShow(false)
        }
    }, [current])

    const openNotification = (title, mess) => {
        notification.open({
            message: title,
            description: mess
        });
    };

    const identifiedBestReturns = () => {
        setTimeout(() => {
            setShow(true)
            openNotification('Pool found', 'Pick Pool #1')
            wrapSetter(3)
        }, 2000)
    }


    return (
        <>
            <OverlayTrigger
                show={show}
                key='bottom'
                placement='bottom'
                overlay={
                    <Tooltip>
                        <strong>Found !</strong>
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
                    Identify Best Returns
                </div>
            </OverlayTrigger>
        </>

    )
}

export default Screen2