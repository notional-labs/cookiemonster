import { useState, useEffect } from 'react'
import { getBalance } from '../../helpers/getBalance';
import { Image, Table, Typography, message, } from 'antd';
import './Asset.css'

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

const Asset = ({ address }) => {
    const [listAsset, setListAsset] = useState([]);

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
        console.log(address)
        if (address !== '') {
            (() => {
                getBalance(undefined, address).then(balances => {
                    if (balances.length > 0) {
                        setListAsset([...balances])
                        success()
                    }
                    warning()
                }).catch(() => {
                    error()
                })
            })()
        }
        else {
            (() => {
                getBalance(undefined, 'osmo1cy2fkq04yh7zm6v52dm525pvx0fph7ed75lnz7').then(balances => {
                    if (balances.length > 0) {
                        setListAsset([...balances])
                        success()
                    }
                }).catch(() => {
                    error()
                })
            })()
        }
    }, [])

    const columns = [
        {
            dataIndex: 'logo',
            key: 'logo',
            fixed: 'left',
            render: (logo) => {
                <Image
                    width={50}
                    src={logo}
                />
            }
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
    return (
        <div style={style.div}>
            <div style={style.divTitle}>
                <Title>Osmosis Assets</Title>
            </div>
            <hr />
            <Table
                dataSource={listAsset}
                columns={columns}
                style={{ marginTop: '3rem', borderRadius: '5px' }}
            />
        </div>
    )
}

export default Asset;