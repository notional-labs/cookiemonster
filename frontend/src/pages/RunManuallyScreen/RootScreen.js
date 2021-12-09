import { useState, useCallback } from 'react'
import { Typography, message } from 'antd';
import { CaretRightFilled, RightOutlined } from '@ant-design/icons'
import Screen1 from './Screen1';
import Screen2 from './Screen2';
import Screen3 from './Screen3';
import Screen4 from './Screen4';

const { Title, } = Typography;

const style = {
    flexDiv: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    divTitle: {
        display: 'flex',
        justifyContent: 'left',
        alignContent: 'left',
        textAlign: 'left',
    },
    div: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        alignContent: 'center'
    },
    button: {
        borderWidth: '0px',
        borderRadius: '10px',
        color: '#ffffff',
        boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.5), 0 6px 20px 0 rgba(0, 0, 0, 0.5)',
        width: "30%",
        height: "5rem",
        padding: '2rem',
        paddingTop: '1rem',
        marginTop: '10rem',
    },
    buttonText: {
        fontSize: '30px',
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'center',
    }
}

const RootScreen = ({ cookieMonster, account }) => {
    const [current, setCurrent] = useState(0)

    const wrapSetter = useCallback(value => {
        setCurrent(value)
    }, [setCurrent])

    const handleEnter = (e) => {
        e.target.style.transform = 'scale(1.01)'
    }

    const handleLeave = (e) => {
        e.target.style.transform = 'scale(1)'
    }

    const startManual = () => {
        if(cookieMonster !== '') {
            setCurrent(1)
        }
        else{
            warning()
        }
    }

    const warning = () => {
        message.warning('Connect to BeanStalk to use this feature', 2);
    };

    return (
        <div style={style.div}>
            <div style={style.divTitle}>
                <Title>BeanStalk</Title>
            </div>
            <hr />
            <div style={{ ...style.flexDiv, marginTop: '2rem' }}>
                <Screen1 current={current} wrapSetter={wrapSetter} cookieMonster={cookieMonster} account={account} />
                <CaretRightFilled style={{ fontSize: '10rem', color: '#7a7a7a' }} />
                <Screen2 current={current} wrapSetter={wrapSetter} account={account} />
                <CaretRightFilled style={{ fontSize: '10rem', color: '#7a7a7a' }} />
                <Screen3 current={current} wrapSetter={wrapSetter} account={account} />
                <CaretRightFilled style={{ fontSize: '10rem', color: '#7a7a7a' }} />
                <Screen4 current={current} wrapSetter={wrapSetter} account={account} />
            </div>
            <div>
                <button 
                    style={{ ...style.button, backgroundColor: '#8abf80', }}
                    onClick={startManual}
                    onMouseEnter={handleEnter}
                    onMouseLeave={handleLeave}>
                    <div style={style.buttonText}>
                        <p>Run Manually</p> <RightOutlined style={{ marginTop: '0.6rem' }} />
                    </div>
                </button>
            </div>
        </div>
    )
}

export default RootScreen