# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: trader.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='trader.proto',
  package='trader',
  syntax='proto3',
  serialized_options=b'Z\010.;trader',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0ctrader.proto\x12\x06trader\"G\n\x18\x43reateMarketOrderRequest\x12\x0e\n\x06opType\x18\x01 \x01(\t\x12\x0e\n\x06ticker\x18\x02 \x01(\t\x12\x0b\n\x03qty\x18\x03 \x01(\x05\"\x1b\n\x19\x43reateMarketOrderResponse\"U\n\x17\x43reateLimitOrderRequest\x12\x0e\n\x06opType\x18\x01 \x01(\t\x12\x0e\n\x06ticker\x18\x02 \x01(\t\x12\r\n\x05price\x18\x03 \x01(\x02\x12\x0b\n\x03qty\x18\x04 \x01(\x05\"\x1a\n\x18\x43reateLimitOrderResponse2\xbd\x01\n\x06Trader\x12Z\n\x11\x43reateMarketOrder\x12 .trader.CreateMarketOrderRequest\x1a!.trader.CreateMarketOrderResponse\"\x00\x12W\n\x10\x43reateLimitOrder\x12\x1f.trader.CreateLimitOrderRequest\x1a .trader.CreateLimitOrderResponse\"\x00\x42\nZ\x08.;traderb\x06proto3'
)




_CREATEMARKETORDERREQUEST = _descriptor.Descriptor(
  name='CreateMarketOrderRequest',
  full_name='trader.CreateMarketOrderRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='opType', full_name='trader.CreateMarketOrderRequest.opType', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ticker', full_name='trader.CreateMarketOrderRequest.ticker', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='qty', full_name='trader.CreateMarketOrderRequest.qty', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=95,
)


_CREATEMARKETORDERRESPONSE = _descriptor.Descriptor(
  name='CreateMarketOrderResponse',
  full_name='trader.CreateMarketOrderResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=97,
  serialized_end=124,
)


_CREATELIMITORDERREQUEST = _descriptor.Descriptor(
  name='CreateLimitOrderRequest',
  full_name='trader.CreateLimitOrderRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='opType', full_name='trader.CreateLimitOrderRequest.opType', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ticker', full_name='trader.CreateLimitOrderRequest.ticker', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='price', full_name='trader.CreateLimitOrderRequest.price', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='qty', full_name='trader.CreateLimitOrderRequest.qty', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=126,
  serialized_end=211,
)


_CREATELIMITORDERRESPONSE = _descriptor.Descriptor(
  name='CreateLimitOrderResponse',
  full_name='trader.CreateLimitOrderResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=213,
  serialized_end=239,
)

DESCRIPTOR.message_types_by_name['CreateMarketOrderRequest'] = _CREATEMARKETORDERREQUEST
DESCRIPTOR.message_types_by_name['CreateMarketOrderResponse'] = _CREATEMARKETORDERRESPONSE
DESCRIPTOR.message_types_by_name['CreateLimitOrderRequest'] = _CREATELIMITORDERREQUEST
DESCRIPTOR.message_types_by_name['CreateLimitOrderResponse'] = _CREATELIMITORDERRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateMarketOrderRequest = _reflection.GeneratedProtocolMessageType('CreateMarketOrderRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEMARKETORDERREQUEST,
  '__module__' : 'trader_pb2'
  # @@protoc_insertion_point(class_scope:trader.CreateMarketOrderRequest)
  })
_sym_db.RegisterMessage(CreateMarketOrderRequest)

CreateMarketOrderResponse = _reflection.GeneratedProtocolMessageType('CreateMarketOrderResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEMARKETORDERRESPONSE,
  '__module__' : 'trader_pb2'
  # @@protoc_insertion_point(class_scope:trader.CreateMarketOrderResponse)
  })
_sym_db.RegisterMessage(CreateMarketOrderResponse)

CreateLimitOrderRequest = _reflection.GeneratedProtocolMessageType('CreateLimitOrderRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATELIMITORDERREQUEST,
  '__module__' : 'trader_pb2'
  # @@protoc_insertion_point(class_scope:trader.CreateLimitOrderRequest)
  })
_sym_db.RegisterMessage(CreateLimitOrderRequest)

CreateLimitOrderResponse = _reflection.GeneratedProtocolMessageType('CreateLimitOrderResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATELIMITORDERRESPONSE,
  '__module__' : 'trader_pb2'
  # @@protoc_insertion_point(class_scope:trader.CreateLimitOrderResponse)
  })
_sym_db.RegisterMessage(CreateLimitOrderResponse)


DESCRIPTOR._options = None

_TRADER = _descriptor.ServiceDescriptor(
  name='Trader',
  full_name='trader.Trader',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=242,
  serialized_end=431,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateMarketOrder',
    full_name='trader.Trader.CreateMarketOrder',
    index=0,
    containing_service=None,
    input_type=_CREATEMARKETORDERREQUEST,
    output_type=_CREATEMARKETORDERRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateLimitOrder',
    full_name='trader.Trader.CreateLimitOrder',
    index=1,
    containing_service=None,
    input_type=_CREATELIMITORDERREQUEST,
    output_type=_CREATELIMITORDERRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TRADER)

DESCRIPTOR.services_by_name['Trader'] = _TRADER

# @@protoc_insertion_point(module_scope)