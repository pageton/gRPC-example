# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: proto/service.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'proto/service.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13proto/service.proto\x12\x07service\"%\n\x07Request\x12\x0c\n\x04path\x18\x01 \x01(\t\x12\x0c\n\x04size\x18\x02 \x01(\x05\"6\n\x08Response\x12\x0e\n\x06status\x18\x01 \x01(\x08\x12\x0c\n\x04path\x18\x02 \x01(\t\x12\x0c\n\x04size\x18\x03 \x01(\x05\x32\x45\n\x0eImageProcessor\x12\x33\n\x0cProcessImage\x12\x10.service.Request\x1a\x11.service.ResponseB\x03Z\x01.b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\001.'
  _globals['_REQUEST']._serialized_start=32
  _globals['_REQUEST']._serialized_end=69
  _globals['_RESPONSE']._serialized_start=71
  _globals['_RESPONSE']._serialized_end=125
  _globals['_IMAGEPROCESSOR']._serialized_start=127
  _globals['_IMAGEPROCESSOR']._serialized_end=196
# @@protoc_insertion_point(module_scope)
