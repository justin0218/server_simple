// source: blogView.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.blog_view', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.blog_view = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.blog_view, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.blog_view.displayName = 'proto.blog_view';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.blog_view.prototype.toObject = function(opt_includeInstance) {
  return proto.blog_view.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blog_view} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog_view.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    country: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    city: jspb.Message.getFieldWithDefault(msg, 4, ""),
    isp: jspb.Message.getFieldWithDefault(msg, 5, ""),
    blogId: jspb.Message.getFieldWithDefault(msg, 6, 0),
    ip: jspb.Message.getFieldWithDefault(msg, 7, ""),
    createTime: jspb.Message.getFieldWithDefault(msg, 8, ""),
    updateTime: jspb.Message.getFieldWithDefault(msg, 9, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.blog_view}
 */
proto.blog_view.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blog_view;
  return proto.blog_view.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blog_view} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blog_view}
 */
proto.blog_view.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCountry(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCity(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setIsp(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setBlogId(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setIp(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreateTime(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setUpdateTime(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.blog_view.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blog_view.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blog_view} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blog_view.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getCountry();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCity();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getIsp();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getBlogId();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getIp();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getCreateTime();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getUpdateTime();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional int32 id = 1;
 * @return {number}
 */
proto.blog_view.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string country = 2;
 * @return {string}
 */
proto.blog_view.prototype.getCountry = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setCountry = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.blog_view.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string city = 4;
 * @return {string}
 */
proto.blog_view.prototype.getCity = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setCity = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string isp = 5;
 * @return {string}
 */
proto.blog_view.prototype.getIsp = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setIsp = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional int32 blog_id = 6;
 * @return {number}
 */
proto.blog_view.prototype.getBlogId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setBlogId = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional string ip = 7;
 * @return {string}
 */
proto.blog_view.prototype.getIp = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setIp = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string create_time = 8;
 * @return {string}
 */
proto.blog_view.prototype.getCreateTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setCreateTime = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string update_time = 9;
 * @return {string}
 */
proto.blog_view.prototype.getUpdateTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.blog_view} returns this
 */
proto.blog_view.prototype.setUpdateTime = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


goog.object.extend(exports, proto);
