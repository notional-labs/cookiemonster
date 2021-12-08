import { TransactionOutlined } from '@ant-design/icons';

const style = {
    button: {
        borderWidth: '0px',
        borderRadius: '5px',
        size: '10em',
        backgroundColor: '#9c8bad',
        color: '#ffffff',
        width: "80%",
        height: "3rem",
    },
}

const DepositButton = ({ collapsed, wrapSetter }) => {
    const handleClick = () => {
        wrapSetter(true)
    }

    return (
        <div style={{ marginTop: '34rem', marginBottom: '0.3rem' }}>
            <hr />
            <button style={{ ...style.button, fontSize: !collapsed ? '15px' : '10px' }}
                onClick={handleClick}>
                {!collapsed ? <div>Deposit <TransactionOutlined style={{ fontSize: '1.5rem' }} /></div>
                    :
                    (<TransactionOutlined style={{ fontSize: '1.5rem' }} />)}
            </button>
        </div>
    )
}

export default DepositButton