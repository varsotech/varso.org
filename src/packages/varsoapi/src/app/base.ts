/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "varso_app";

export interface Organization {
  uuid: string;
  uniqueName: string;
  name: string;
  description: string;
  websiteUrl: string;
}

export interface RSSFeed {
  items: RSSItem[];
}

export interface RSSItem {
  title: string;
  description: string;
  content: string;
  link: string;
  links: string[];
  updateDate: Date | undefined;
  publishDate: Date | undefined;
  authors: RSSAuthor[];
  guid: string;
  image: RSSImage | undefined;
  categories: string[];
  organizationUuid: string;
  id: string;
}

export interface RSSAuthor {
  email: string;
  name: string;
}

export interface RSSImage {
  url: string;
  title: string;
  blur: boolean;
}

function createBaseOrganization(): Organization {
  return { uuid: "", uniqueName: "", name: "", description: "", websiteUrl: "" };
}

export const Organization = {
  encode(message: Organization, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.uniqueName !== "") {
      writer.uint32(18).string(message.uniqueName);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.websiteUrl !== "") {
      writer.uint32(42).string(message.websiteUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Organization {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOrganization();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uuid = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.uniqueName = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.description = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.websiteUrl = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Organization {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      uniqueName: isSet(object.uniqueName) ? String(object.uniqueName) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      websiteUrl: isSet(object.websiteUrl) ? String(object.websiteUrl) : "",
    };
  },

  toJSON(message: Organization): unknown {
    const obj: any = {};
    if (message.uuid !== "") {
      obj.uuid = message.uuid;
    }
    if (message.uniqueName !== "") {
      obj.uniqueName = message.uniqueName;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.websiteUrl !== "") {
      obj.websiteUrl = message.websiteUrl;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Organization>, I>>(base?: I): Organization {
    return Organization.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Organization>, I>>(object: I): Organization {
    const message = createBaseOrganization();
    message.uuid = object.uuid ?? "";
    message.uniqueName = object.uniqueName ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.websiteUrl = object.websiteUrl ?? "";
    return message;
  },
};

function createBaseRSSFeed(): RSSFeed {
  return { items: [] };
}

export const RSSFeed = {
  encode(message: RSSFeed, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.items) {
      RSSItem.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RSSFeed {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRSSFeed();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.items.push(RSSItem.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RSSFeed {
    return { items: Array.isArray(object?.items) ? object.items.map((e: any) => RSSItem.fromJSON(e)) : [] };
  },

  toJSON(message: RSSFeed): unknown {
    const obj: any = {};
    if (message.items?.length) {
      obj.items = message.items.map((e) => RSSItem.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RSSFeed>, I>>(base?: I): RSSFeed {
    return RSSFeed.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RSSFeed>, I>>(object: I): RSSFeed {
    const message = createBaseRSSFeed();
    message.items = object.items?.map((e) => RSSItem.fromPartial(e)) || [];
    return message;
  },
};

function createBaseRSSItem(): RSSItem {
  return {
    title: "",
    description: "",
    content: "",
    link: "",
    links: [],
    updateDate: undefined,
    publishDate: undefined,
    authors: [],
    guid: "",
    image: undefined,
    categories: [],
    organizationUuid: "",
    id: "",
  };
}

export const RSSItem = {
  encode(message: RSSItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.content !== "") {
      writer.uint32(26).string(message.content);
    }
    if (message.link !== "") {
      writer.uint32(34).string(message.link);
    }
    for (const v of message.links) {
      writer.uint32(42).string(v!);
    }
    if (message.updateDate !== undefined) {
      Timestamp.encode(toTimestamp(message.updateDate), writer.uint32(50).fork()).ldelim();
    }
    if (message.publishDate !== undefined) {
      Timestamp.encode(toTimestamp(message.publishDate), writer.uint32(58).fork()).ldelim();
    }
    for (const v of message.authors) {
      RSSAuthor.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.guid !== "") {
      writer.uint32(74).string(message.guid);
    }
    if (message.image !== undefined) {
      RSSImage.encode(message.image, writer.uint32(82).fork()).ldelim();
    }
    for (const v of message.categories) {
      writer.uint32(90).string(v!);
    }
    if (message.organizationUuid !== "") {
      writer.uint32(98).string(message.organizationUuid);
    }
    if (message.id !== "") {
      writer.uint32(106).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RSSItem {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRSSItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.description = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.content = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.link = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.links.push(reader.string());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.updateDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.publishDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.authors.push(RSSAuthor.decode(reader, reader.uint32()));
          continue;
        case 9:
          if (tag !== 74) {
            break;
          }

          message.guid = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.image = RSSImage.decode(reader, reader.uint32());
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.categories.push(reader.string());
          continue;
        case 12:
          if (tag !== 98) {
            break;
          }

          message.organizationUuid = reader.string();
          continue;
        case 13:
          if (tag !== 106) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RSSItem {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      content: isSet(object.content) ? String(object.content) : "",
      link: isSet(object.link) ? String(object.link) : "",
      links: Array.isArray(object?.links) ? object.links.map((e: any) => String(e)) : [],
      updateDate: isSet(object.updateDate) ? fromJsonTimestamp(object.updateDate) : undefined,
      publishDate: isSet(object.publishDate) ? fromJsonTimestamp(object.publishDate) : undefined,
      authors: Array.isArray(object?.authors) ? object.authors.map((e: any) => RSSAuthor.fromJSON(e)) : [],
      guid: isSet(object.guid) ? String(object.guid) : "",
      image: isSet(object.image) ? RSSImage.fromJSON(object.image) : undefined,
      categories: Array.isArray(object?.categories) ? object.categories.map((e: any) => String(e)) : [],
      organizationUuid: isSet(object.organizationUuid) ? String(object.organizationUuid) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: RSSItem): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.content !== "") {
      obj.content = message.content;
    }
    if (message.link !== "") {
      obj.link = message.link;
    }
    if (message.links?.length) {
      obj.links = message.links;
    }
    if (message.updateDate !== undefined) {
      obj.updateDate = message.updateDate.toISOString();
    }
    if (message.publishDate !== undefined) {
      obj.publishDate = message.publishDate.toISOString();
    }
    if (message.authors?.length) {
      obj.authors = message.authors.map((e) => RSSAuthor.toJSON(e));
    }
    if (message.guid !== "") {
      obj.guid = message.guid;
    }
    if (message.image !== undefined) {
      obj.image = RSSImage.toJSON(message.image);
    }
    if (message.categories?.length) {
      obj.categories = message.categories;
    }
    if (message.organizationUuid !== "") {
      obj.organizationUuid = message.organizationUuid;
    }
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RSSItem>, I>>(base?: I): RSSItem {
    return RSSItem.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RSSItem>, I>>(object: I): RSSItem {
    const message = createBaseRSSItem();
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.content = object.content ?? "";
    message.link = object.link ?? "";
    message.links = object.links?.map((e) => e) || [];
    message.updateDate = object.updateDate ?? undefined;
    message.publishDate = object.publishDate ?? undefined;
    message.authors = object.authors?.map((e) => RSSAuthor.fromPartial(e)) || [];
    message.guid = object.guid ?? "";
    message.image = (object.image !== undefined && object.image !== null)
      ? RSSImage.fromPartial(object.image)
      : undefined;
    message.categories = object.categories?.map((e) => e) || [];
    message.organizationUuid = object.organizationUuid ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseRSSAuthor(): RSSAuthor {
  return { email: "", name: "" };
}

export const RSSAuthor = {
  encode(message: RSSAuthor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RSSAuthor {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRSSAuthor();
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

  fromJSON(object: any): RSSAuthor {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      name: isSet(object.name) ? String(object.name) : "",
    };
  },

  toJSON(message: RSSAuthor): unknown {
    const obj: any = {};
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RSSAuthor>, I>>(base?: I): RSSAuthor {
    return RSSAuthor.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RSSAuthor>, I>>(object: I): RSSAuthor {
    const message = createBaseRSSAuthor();
    message.email = object.email ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseRSSImage(): RSSImage {
  return { url: "", title: "", blur: false };
}

export const RSSImage = {
  encode(message: RSSImage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.blur === true) {
      writer.uint32(24).bool(message.blur);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RSSImage {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRSSImage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.url = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.blur = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RSSImage {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      title: isSet(object.title) ? String(object.title) : "",
      blur: isSet(object.blur) ? Boolean(object.blur) : false,
    };
  },

  toJSON(message: RSSImage): unknown {
    const obj: any = {};
    if (message.url !== "") {
      obj.url = message.url;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.blur === true) {
      obj.blur = message.blur;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RSSImage>, I>>(base?: I): RSSImage {
    return RSSImage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RSSImage>, I>>(object: I): RSSImage {
    const message = createBaseRSSImage();
    message.url = object.url ?? "";
    message.title = object.title ?? "";
    message.blur = object.blur ?? false;
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
