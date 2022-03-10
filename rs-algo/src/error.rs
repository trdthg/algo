use std::fmt::Display;

pub enum Error {
    StrErr { msg: String },
}

impl Display for Error {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::StrErr { msg } => write!(f, "{msg}"),
        }
    }
}

pub type Result<T> = std::result::Result<T, Error>;
