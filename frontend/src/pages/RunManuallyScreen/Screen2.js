import { useState, useEffect} from 'react'

const Screen2 = ({ current, wrapSetter }) => {
    const [state, setState] = useState('pending')

    useEffect(() => {
        if (current === 2) {
            setState('running')
        }
        else if( current === 0){
            setState('pending')
        }
    }, [current])

    const identifiedBestReturns = () => {
        setTimeout(() => {
            wrapSetter(3)
        }, 2000)
    }

    if(state === 'running'){
        identifiedBestReturns()
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
            Identify Best Returns 
        </div>
    )
}

export default Screen2