import { useState, useEffect } from 'react'
import { getBalance } from '../../helpers/getBalance';
import { getTotal } from '../../helpers/convertPrice';
import { Image, Table, Typography, message, } from 'antd';
import './Asset.css'
import image from '../../assets/img/plant.png'

const { Title, } = Typography;

const style = {
    divTitle: {
        display: 'flex',
        justifyContent: 'left',
        alignContent: 'left',
        textAlign: 'left',
    },
    div: {
        marginBottom: '8rem',
        borderRadius: '10px'
    }
}

const Asset = ({ cookieMonster }) => {
    const [listAsset, setListAsset] = useState([]);
    const [asset, setAsset] = useState(0)

    const success = () => {
        message.success('Fetching success', 1);
    };

    const error = () => {
        message.error('Fetching failed', 1);
    };

    const warning = () => {
        message.warning('No Assets Yet', 1);
    };

    useEffect(() => {
        console.log(cookieMonster)
        if (cookieMonster !== '') {
            (() => {
                getBalance(undefined, cookieMonster).then(balances => {
                    if (balances.length > 0) {
                        success()
                        setListAsset([...balances])
                        getTotal(balances).then(total => setAsset(total))
                    }
                    else{
                        warning()
                    }
                }).catch(() => {
                    error()
                })
            })()
        }
    }, [cookieMonster])

    const columns = [
        {
            dataIndex: 'logo',
            key: 'logo',
            fixed: 'left',
            width: 20,
            render: (logo) => <img width={40} src={logo} />
            
        },
        {
            title: 'Asset/Chain',
            dataIndex: 'name',
            key: 'name',
            fixed: 'left',

        },
        {
            title: 'Amount',
            dataIndex: 'amount',
            key: 'amount',
            fixed: 'left',
        },
    ];

    return cookieMonster !== '' ? (
        <div style={style.div}>
            <div style={style.divTitle}>
                <Title>BeanStalk Assets</Title>
            </div>
            <hr />
            <div style={{marginBottom: '1rem', ...style.divTitle}}>
                Total Assets: {asset} USD
            </div>
            <Table
                dataSource={listAsset}
                columns={columns}
                style={{ marginTop: '3rem', borderRadius: '5px' }}
            />
        </div>
    ) : (
        <div>
            <img width={400} src={image} style={{opacity: '50%', marginBottom: '3rem', marginTop: '50px'}}/>
            <p style={{fontSize: '30px'}}>Connect To BeanStalk To Check Assets</p>
        </div>
    )
}

export default Asset;