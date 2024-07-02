/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { User } from "./base";

export const protobufPackage = "varso_auth";

export interface BanUserRequest {
}

export interface BanUserResponse {
  bannedUserDiscordId: string;
}

export interface DiscordLoginRequest {
  internalLoginAuthSecret: string;
  discordUserId: string;
}

export interface DiscordLoginResponse {
  token: string;
}

export interface GetUserResponse {
  user: User | undefined;
}

export interface GetUserByDiscordIdResponse {
  user: User | undefined;
}

export interface InternalLoginRequest {
  internalLoginAuthSecret: string;
}

export interface InternalLoginResponse {
  token: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
}

export interface RegisterResponse {
  token: string;
}

export interface EasyRegisterRequest {
  name: string;
}

export interface EasyRegisterResponse {
  token: string;
}

export interface LoginRequest {
  usernameOrEmail: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface UnbanUserRequest {
}

export interface UnbanUserResponse {
  userDiscordId: string;
}

function createBaseBanUserRequest(): BanUserRequest {
  return {};
}

export const BanUserRequest = {
  encode(_: BanUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BanUserRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBanUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): BanUserRequest {
    return {};
  },

  toJSON(_: BanUserRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<BanUserRequest>, I>>(base?: I): BanUserRequest {
    return BanUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<BanUserRequest>, I>>(_: I): BanUserRequest {
    const message = createBaseBanUserRequest();
    return message;
  },
};

function createBaseBanUserResponse(): BanUserResponse {
  return { bannedUserDiscordId: "" };
}

export const BanUserResponse = {
  encode(message: BanUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bannedUserDiscordId !== "") {
      writer.uint32(10).string(message.bannedUserDiscordId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BanUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBanUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.bannedUserDiscordId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BanUserResponse {
    return { bannedUserDiscordId: isSet(object.bannedUserDiscordId) ? String(object.bannedUserDiscordId) : "" };
  },

  toJSON(message: BanUserResponse): unknown {
    const obj: any = {};
    if (message.bannedUserDiscordId !== "") {
      obj.bannedUserDiscordId = message.bannedUserDiscordId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<BanUserResponse>, I>>(base?: I): BanUserResponse {
    return BanUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<BanUserResponse>, I>>(object: I): BanUserResponse {
    const message = createBaseBanUserResponse();
    message.bannedUserDiscordId = object.bannedUserDiscordId ?? "";
    return message;
  },
};

function createBaseDiscordLoginRequest(): DiscordLoginRequest {
  return { internalLoginAuthSecret: "", discordUserId: "" };
}

export const DiscordLoginRequest = {
  encode(message: DiscordLoginRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.internalLoginAuthSecret !== "") {
      writer.uint32(10).string(message.internalLoginAuthSecret);
    }
    if (message.discordUserId !== "") {
      writer.uint32(18).string(message.discordUserId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DiscordLoginRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDiscordLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.internalLoginAuthSecret = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.discordUserId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DiscordLoginRequest {
    return {
      internalLoginAuthSecret: isSet(object.internalLoginAuthSecret) ? String(object.internalLoginAuthSecret) : "",
      discordUserId: isSet(object.discordUserId) ? String(object.discordUserId) : "",
    };
  },

  toJSON(message: DiscordLoginRequest): unknown {
    const obj: any = {};
    if (message.internalLoginAuthSecret !== "") {
      obj.internalLoginAuthSecret = message.internalLoginAuthSecret;
    }
    if (message.discordUserId !== "") {
      obj.discordUserId = message.discordUserId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DiscordLoginRequest>, I>>(base?: I): DiscordLoginRequest {
    return DiscordLoginRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DiscordLoginRequest>, I>>(object: I): DiscordLoginRequest {
    const message = createBaseDiscordLoginRequest();
    message.internalLoginAuthSecret = object.internalLoginAuthSecret ?? "";
    message.discordUserId = object.discordUserId ?? "";
    return message;
  },
};

function createBaseDiscordLoginResponse(): DiscordLoginResponse {
  return { token: "" };
}

export const DiscordLoginResponse = {
  encode(message: DiscordLoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DiscordLoginResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDiscordLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DiscordLoginResponse {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: DiscordLoginResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DiscordLoginResponse>, I>>(base?: I): DiscordLoginResponse {
    return DiscordLoginResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DiscordLoginResponse>, I>>(object: I): DiscordLoginResponse {
    const message = createBaseDiscordLoginResponse();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseGetUserResponse(): GetUserResponse {
  return { user: undefined };
}

export const GetUserResponse = {
  encode(message: GetUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: GetUserResponse): unknown {
    const obj: any = {};
    if (message.user !== undefined) {
      obj.user = User.toJSON(message.user);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserResponse>, I>>(base?: I): GetUserResponse {
    return GetUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserResponse>, I>>(object: I): GetUserResponse {
    const message = createBaseGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseGetUserByDiscordIdResponse(): GetUserByDiscordIdResponse {
  return { user: undefined };
}

export const GetUserByDiscordIdResponse = {
  encode(message: GetUserByDiscordIdResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserByDiscordIdResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserByDiscordIdResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserByDiscordIdResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: GetUserByDiscordIdResponse): unknown {
    const obj: any = {};
    if (message.user !== undefined) {
      obj.user = User.toJSON(message.user);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserByDiscordIdResponse>, I>>(base?: I): GetUserByDiscordIdResponse {
    return GetUserByDiscordIdResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserByDiscordIdResponse>, I>>(object: I): GetUserByDiscordIdResponse {
    const message = createBaseGetUserByDiscordIdResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseInternalLoginRequest(): InternalLoginRequest {
  return { internalLoginAuthSecret: "" };
}

export const InternalLoginRequest = {
  encode(message: InternalLoginRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.internalLoginAuthSecret !== "") {
      writer.uint32(10).string(message.internalLoginAuthSecret);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InternalLoginRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInternalLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.internalLoginAuthSecret = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InternalLoginRequest {
    return {
      internalLoginAuthSecret: isSet(object.internalLoginAuthSecret) ? String(object.internalLoginAuthSecret) : "",
    };
  },

  toJSON(message: InternalLoginRequest): unknown {
    const obj: any = {};
    if (message.internalLoginAuthSecret !== "") {
      obj.internalLoginAuthSecret = message.internalLoginAuthSecret;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<InternalLoginRequest>, I>>(base?: I): InternalLoginRequest {
    return InternalLoginRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<InternalLoginRequest>, I>>(object: I): InternalLoginRequest {
    const message = createBaseInternalLoginRequest();
    message.internalLoginAuthSecret = object.internalLoginAuthSecret ?? "";
    return message;
  },
};

function createBaseInternalLoginResponse(): InternalLoginResponse {
  return { token: "" };
}

export const InternalLoginResponse = {
  encode(message: InternalLoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InternalLoginResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInternalLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InternalLoginResponse {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: InternalLoginResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<InternalLoginResponse>, I>>(base?: I): InternalLoginResponse {
    return InternalLoginResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<InternalLoginResponse>, I>>(object: I): InternalLoginResponse {
    const message = createBaseInternalLoginResponse();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseRegisterRequest(): RegisterRequest {
  return { email: "", password: "" };
}

export const RegisterRequest = {
  encode(message: RegisterRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.email = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.password = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RegisterRequest {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: RegisterRequest): unknown {
    const obj: any = {};
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.password !== "") {
      obj.password = message.password;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RegisterRequest>, I>>(base?: I): RegisterRequest {
    return RegisterRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RegisterRequest>, I>>(object: I): RegisterRequest {
    const message = createBaseRegisterRequest();
    message.email = object.email ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseRegisterResponse(): RegisterResponse {
  return { token: "" };
}

export const RegisterResponse = {
  encode(message: RegisterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RegisterResponse {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: RegisterResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RegisterResponse>, I>>(base?: I): RegisterResponse {
    return RegisterResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RegisterResponse>, I>>(object: I): RegisterResponse {
    const message = createBaseRegisterResponse();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseEasyRegisterRequest(): EasyRegisterRequest {
  return { name: "" };
}

export const EasyRegisterRequest = {
  encode(message: EasyRegisterRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EasyRegisterRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEasyRegisterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EasyRegisterRequest {
    return { name: isSet(object.name) ? String(object.name) : "" };
  },

  toJSON(message: EasyRegisterRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<EasyRegisterRequest>, I>>(base?: I): EasyRegisterRequest {
    return EasyRegisterRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<EasyRegisterRequest>, I>>(object: I): EasyRegisterRequest {
    const message = createBaseEasyRegisterRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseEasyRegisterResponse(): EasyRegisterResponse {
  return { token: "" };
}

export const EasyRegisterResponse = {
  encode(message: EasyRegisterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EasyRegisterResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEasyRegisterResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EasyRegisterResponse {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: EasyRegisterResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<EasyRegisterResponse>, I>>(base?: I): EasyRegisterResponse {
    return EasyRegisterResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<EasyRegisterResponse>, I>>(object: I): EasyRegisterResponse {
    const message = createBaseEasyRegisterResponse();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseLoginRequest(): LoginRequest {
  return { usernameOrEmail: "", password: "" };
}

export const LoginRequest = {
  encode(message: LoginRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.usernameOrEmail !== "") {
      writer.uint32(10).string(message.usernameOrEmail);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.usernameOrEmail = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.password = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LoginRequest {
    return {
      usernameOrEmail: isSet(object.usernameOrEmail) ? String(object.usernameOrEmail) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: LoginRequest): unknown {
    const obj: any = {};
    if (message.usernameOrEmail !== "") {
      obj.usernameOrEmail = message.usernameOrEmail;
    }
    if (message.password !== "") {
      obj.password = message.password;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<LoginRequest>, I>>(base?: I): LoginRequest {
    return LoginRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<LoginRequest>, I>>(object: I): LoginRequest {
    const message = createBaseLoginRequest();
    message.usernameOrEmail = object.usernameOrEmail ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseLoginResponse(): LoginResponse {
  return { token: "" };
}

export const LoginResponse = {
  encode(message: LoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LoginResponse {
    return { token: isSet(object.token) ? String(object.token) : "" };
  },

  toJSON(message: LoginResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<LoginResponse>, I>>(base?: I): LoginResponse {
    return LoginResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<LoginResponse>, I>>(object: I): LoginResponse {
    const message = createBaseLoginResponse();
    message.token = object.token ?? "";
    return message;
  },
};

function createBaseUnbanUserRequest(): UnbanUserRequest {
  return {};
}

export const UnbanUserRequest = {
  encode(_: UnbanUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnbanUserRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnbanUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): UnbanUserRequest {
    return {};
  },

  toJSON(_: UnbanUserRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<UnbanUserRequest>, I>>(base?: I): UnbanUserRequest {
    return UnbanUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UnbanUserRequest>, I>>(_: I): UnbanUserRequest {
    const message = createBaseUnbanUserRequest();
    return message;
  },
};

function createBaseUnbanUserResponse(): UnbanUserResponse {
  return { userDiscordId: "" };
}

export const UnbanUserResponse = {
  encode(message: UnbanUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userDiscordId !== "") {
      writer.uint32(10).string(message.userDiscordId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnbanUserResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnbanUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.userDiscordId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UnbanUserResponse {
    return { userDiscordId: isSet(object.userDiscordId) ? String(object.userDiscordId) : "" };
  },

  toJSON(message: UnbanUserResponse): unknown {
    const obj: any = {};
    if (message.userDiscordId !== "") {
      obj.userDiscordId = message.userDiscordId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UnbanUserResponse>, I>>(base?: I): UnbanUserResponse {
    return UnbanUserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UnbanUserResponse>, I>>(object: I): UnbanUserResponse {
    const message = createBaseUnbanUserResponse();
    message.userDiscordId = object.userDiscordId ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
