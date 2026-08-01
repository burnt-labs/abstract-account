use cosmos_sdk_proto::{cosmos, traits::TypeUrl};
use cosmwasm_std::CosmosMsg;
use prost::Message;

#[derive(Clone, PartialEq, prost::Message)]
pub struct MsgRegisterAccount {
    #[prost(string, tag = "1")]
    pub sender: String,

    #[prost(bytes = "vec", tag = "3")]
    pub msg: Vec<u8>,

    #[prost(message, repeated, tag = "4")]
    pub funds: Vec<cosmos::base::v1beta1::Coin>,

    #[prost(bytes = "vec", tag = "5")]
    pub salt: Vec<u8>,
}

impl From<MsgRegisterAccount> for CosmosMsg {
    fn from(msg: MsgRegisterAccount) -> Self {
        CosmosMsg::Stargate {
            type_url: MsgRegisterAccount::TYPE_URL.into(),
            value:    msg.encode_to_vec().into(),
        }
    }
}

impl TypeUrl for MsgRegisterAccount {
    const TYPE_URL: &'static str = "/abstractaccount.v1.MsgRegisterAccount";
}

#[allow(dead_code)]
#[derive(Clone, PartialEq, prost::Message)]
pub struct MsgRegisterAccountResponse {
    #[prost(string, tag = "1")]
    pub address: String,

    #[prost(bytes = "vec", tag = "2")]
    pub data: Vec<u8>,
}

impl TypeUrl for MsgRegisterAccountResponse {
    const TYPE_URL: &'static str = "/abstractaccount.v1.MsgRegisterAccountResponse";
}

#[derive(Clone, PartialEq, prost::Message)]
pub struct MsgMigrateAccount {
    #[prost(string, tag = "1")]
    pub sender: String,
}

impl From<MsgMigrateAccount> for CosmosMsg {
    fn from(msg: MsgMigrateAccount) -> Self {
        CosmosMsg::Stargate {
            type_url: MsgMigrateAccount::TYPE_URL.into(),
            value:    msg.encode_to_vec().into(),
        }
    }
}

impl TypeUrl for MsgMigrateAccount {
    const TYPE_URL: &'static str = "/abstractaccount.v1.MsgMigrateAccount";
}

#[allow(dead_code)]
#[derive(Clone, PartialEq, prost::Message)]
pub struct MsgMigrateAccountResponse {
    #[prost(bool, tag = "1")]
    pub migrated: bool,

    #[prost(bytes = "vec", tag = "2")]
    pub data: Vec<u8>,
}

impl TypeUrl for MsgMigrateAccountResponse {
    const TYPE_URL: &'static str = "/abstractaccount.v1.MsgMigrateAccountResponse";
}

// TODO: add definitions for AbstractAccount and NilPubKey
// TODO: add tests
