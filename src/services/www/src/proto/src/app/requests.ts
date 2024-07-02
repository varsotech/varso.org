/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Organization, RSSFeed, RSSItem } from "./base";

export const protobufPackage = "varso_app";

export interface GetOrganizationsResponse {
  organizations: Organization[];
}

export interface GetNewsResponse {
  organizations: { [key: string]: Organization };
  featured: RSSItem | undefined;
  latest: RSSFeed | undefined;
}

export interface GetNewsResponse_OrganizationsEntry {
  key: string;
  value: Organization | undefined;
}

export interface BlurToggleRequest {
  rssItemId: string;
}

function createBaseGetOrganizationsResponse(): GetOrganizationsResponse {
  return { organizations: [] };
}

export const GetOrganizationsResponse = {
  encode(message: GetOrganizationsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.organizations) {
      Organization.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetOrganizationsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetOrganizationsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.organizations.push(Organization.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetOrganizationsResponse {
    return {
      organizations: Array.isArray(object?.organizations)
        ? object.organizations.map((e: any) => Organization.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GetOrganizationsResponse): unknown {
    const obj: any = {};
    if (message.organizations?.length) {
      obj.organizations = message.organizations.map((e) => Organization.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetOrganizationsResponse>, I>>(base?: I): GetOrganizationsResponse {
    return GetOrganizationsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetOrganizationsResponse>, I>>(object: I): GetOrganizationsResponse {
    const message = createBaseGetOrganizationsResponse();
    message.organizations = object.organizations?.map((e) => Organization.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetNewsResponse(): GetNewsResponse {
  return { organizations: {}, featured: undefined, latest: undefined };
}

export const GetNewsResponse = {
  encode(message: GetNewsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.organizations).forEach(([key, value]) => {
      GetNewsResponse_OrganizationsEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    if (message.featured !== undefined) {
      RSSItem.encode(message.featured, writer.uint32(18).fork()).ldelim();
    }
    if (message.latest !== undefined) {
      RSSFeed.encode(message.latest, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          const entry1 = GetNewsResponse_OrganizationsEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.organizations[entry1.key] = entry1.value;
          }
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.featured = RSSItem.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.latest = RSSFeed.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetNewsResponse {
    return {
      organizations: isObject(object.organizations)
        ? Object.entries(object.organizations).reduce<{ [key: string]: Organization }>((acc, [key, value]) => {
          acc[key] = Organization.fromJSON(value);
          return acc;
        }, {})
        : {},
      featured: isSet(object.featured) ? RSSItem.fromJSON(object.featured) : undefined,
      latest: isSet(object.latest) ? RSSFeed.fromJSON(object.latest) : undefined,
    };
  },

  toJSON(message: GetNewsResponse): unknown {
    const obj: any = {};
    if (message.organizations) {
      const entries = Object.entries(message.organizations);
      if (entries.length > 0) {
        obj.organizations = {};
        entries.forEach(([k, v]) => {
          obj.organizations[k] = Organization.toJSON(v);
        });
      }
    }
    if (message.featured !== undefined) {
      obj.featured = RSSItem.toJSON(message.featured);
    }
    if (message.latest !== undefined) {
      obj.latest = RSSFeed.toJSON(message.latest);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetNewsResponse>, I>>(base?: I): GetNewsResponse {
    return GetNewsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetNewsResponse>, I>>(object: I): GetNewsResponse {
    const message = createBaseGetNewsResponse();
    message.organizations = Object.entries(object.organizations ?? {}).reduce<{ [key: string]: Organization }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = Organization.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    message.featured = (object.featured !== undefined && object.featured !== null)
      ? RSSItem.fromPartial(object.featured)
      : undefined;
    message.latest = (object.latest !== undefined && object.latest !== null)
      ? RSSFeed.fromPartial(object.latest)
      : undefined;
    return message;
  },
};

function createBaseGetNewsResponse_OrganizationsEntry(): GetNewsResponse_OrganizationsEntry {
  return { key: "", value: undefined };
}

export const GetNewsResponse_OrganizationsEntry = {
  encode(message: GetNewsResponse_OrganizationsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Organization.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetNewsResponse_OrganizationsEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetNewsResponse_OrganizationsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = Organization.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetNewsResponse_OrganizationsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? Organization.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: GetNewsResponse_OrganizationsEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== undefined) {
      obj.value = Organization.toJSON(message.value);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetNewsResponse_OrganizationsEntry>, I>>(
    base?: I,
  ): GetNewsResponse_OrganizationsEntry {
    return GetNewsResponse_OrganizationsEntry.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetNewsResponse_OrganizationsEntry>, I>>(
    object: I,
  ): GetNewsResponse_OrganizationsEntry {
    const message = createBaseGetNewsResponse_OrganizationsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? Organization.fromPartial(object.value)
      : undefined;
    return message;
  },
};

function createBaseBlurToggleRequest(): BlurToggleRequest {
  return { rssItemId: "" };
}

export const BlurToggleRequest = {
  encode(message: BlurToggleRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.rssItemId !== "") {
      writer.uint32(10).string(message.rssItemId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BlurToggleRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBlurToggleRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.rssItemId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BlurToggleRequest {
    return { rssItemId: isSet(object.rssItemId) ? String(object.rssItemId) : "" };
  },

  toJSON(message: BlurToggleRequest): unknown {
    const obj: any = {};
    if (message.rssItemId !== "") {
      obj.rssItemId = message.rssItemId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<BlurToggleRequest>, I>>(base?: I): BlurToggleRequest {
    return BlurToggleRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<BlurToggleRequest>, I>>(object: I): BlurToggleRequest {
    const message = createBaseBlurToggleRequest();
    message.rssItemId = object.rssItemId ?? "";
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
