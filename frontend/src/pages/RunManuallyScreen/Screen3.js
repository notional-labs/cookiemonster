import { useState, useEffect} from 'react'

const Screen3 = ({ current, wrapSetter }) => {
    const [state, setState] = useState('pending')

    useEffect(() => {
        if (current === 3) {
            setState('running')
        }
        else if( current === 0){
            setState('pending')
        }
    }, [current])

    const autoInvest = () => {
        setTimeout(() => {
            wrapSetter(4)
        }, 2000)
    }

    if(state === 'running'){
        autoInvest()
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
            Auto Investing
        </div>
    )
}

export default Screen3