import { Typography, message } from 'antd';

const { Title, Paragraph } = Typography;

const Account = ({ account }) => {
    return (
        <div>
            <Title level={4}>Wallet info</Title>
            <div>
                <Paragraph copyable={{ text: account.address }}>
                    {account.address.length > 100 ? `${account.address.substring(0, 100)}... ` : `${account.address} `}
                </Paragraph>
            </div>
        </div>
    )
}

export default Account