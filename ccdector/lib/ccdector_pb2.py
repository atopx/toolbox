# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ccdector.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0e\x63\x63\x64\x65\x63tor.proto\x12\x08\x63\x63\x64\x65\x63tor\"\x16\n\x06Params\x12\x0c\n\x04\x66ile\x18\x01 \x01(\x0c\"\xa3\x01\n\x05Reply\x12$\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32\x16.ccdector.Reply.Record\x1at\n\x06Record\x12\x12\n\x05label\x18\x01 \x01(\x05H\x00\x88\x01\x01\x12\x12\n\nconfidence\x18\x02 \x01(\x01\x12\x0c\n\x04left\x18\x03 \x01(\x01\x12\x0b\n\x03top\x18\x04 \x01(\x01\x12\r\n\x05right\x18\x05 \x01(\x01\x12\x0e\n\x06\x62ottom\x18\x06 \x01(\x01\x42\x08\n\x06_label28\n\x06\x44\x65\x63tor\x12.\n\x07Predict\x12\x10.ccdector.Params\x1a\x0f.ccdector.Reply\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ccdector_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_PARAMS']._serialized_start=28
  _globals['_PARAMS']._serialized_end=50
  _globals['_REPLY']._serialized_start=53
  _globals['_REPLY']._serialized_end=216
  _globals['_REPLY_RECORD']._serialized_start=100
  _globals['_REPLY_RECORD']._serialized_end=216
  _globals['_DECTOR']._serialized_start=218
  _globals['_DECTOR']._serialized_end=274
# @@protoc_insertion_point(module_scope)