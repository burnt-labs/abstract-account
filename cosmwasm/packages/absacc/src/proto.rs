use cosmos_sdk_proto::{cosmos};
use cosmwasm_std::{AnyMsg, CosmosMsg, to_json_binary};
use prost::Message;

#[derive(Clone, PartialEq, prost::Message)]
pub struct MsgRegisterAccount {
    #[prost(string, tag = "1")]
    pub sender: String,

    #[prost(uint64, tag = "2")]
    pub code_id: u64,

    #[prost(bytes = "vec", tag = "3")]
    pub msg: Vec<u8>,

    #[prost(message, repeated, tag = "4")]
    pub funds: Vec<cosmos::base::v1beta1::Coin>,

    #[prost(bytes = "vec", tag = "5")]
    pub salt: Vec<u8>,
}

impl From<MsgRegisterAccount> for CosmosMsg {
    fn from(msg: MsgRegisterAccount) -> Self {
        let any_msg: AnyMsg = AnyMsg{
            type_url: String::from("/abstractaccount.v1.MsgRegisterAccount"),
            value:    msg.encode_to_vec().into(),
        };
        CosmosMsg::Any(any_msg)
    }
}

// impl TypeUrl for MsgRegisterAccount {
//     const TYPE_URL: &'static str = "/abstractaccount.v1.MsgRegisterAccount";
// }

#[derive(Clone, PartialEq, prost::Message)]
pub struct MsgRegisterAccountResponse {
    #[prost(string, tag = "1")]
    pub address: String,

    #[prost(bytes = "vec", tag = "2")]
    pub data: Vec<u8>,
}

// impl TypeUrl for MsgRegisterAccountResponse {
//     const TYPE_URL: &'static str = "/abstractaccount.v1.MsgRegisterAccountResponse";
// }

// TODO: add definitions for AbstractAccount and NilPubKey
// TODO: add tests
