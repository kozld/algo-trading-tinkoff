import pytest
from grabber import parse_text

@pytest.fixture
def example_pattern():
    return '(BUY){1} #(\w+)\s?.*\$(\d+\.?\d*)\s?x\s?(\d+)\s?;.*\+(\d+\.?\d*)'

@pytest.fixture
def parse_text_cases():
    return [
        {
            'msg': '❗️BUY #SQ $219.37 x 20; TAKE PROFIT +2.5%',
            'data': [('BUY', 'SQ', '219.37', '20', '2.5')]
        },
        {
            'msg': '❗️BUY #FRHC $51 x 100; TAKE PROFIT +2.5%',
            'data': [('BUY', 'FRHC', '51', '100', '2.5')]
        },
        {
            'msg': '❗BUY #PTON $147.93 x 20; TAKE PROFIT +2.5%',
            'data': [('BUY', 'PTON', '147.93', '20', '2.5')]
        },
        {
            'msg': '❗️BUY #TWTR $45.05 x 80; TAKE PROFIT +2.5%',
            'data': [('BUY', 'TWTR', '45.05', '80', '2.5')]
        }
    ]

@pytest.fixture
def parse_text_cases_ignore():
    return [
        {
            'msg': '✅ SELL #INTU $378.97 x 10 ; Profit: +93$',
            'data': []
        },
        {
            'msg': '✅ SELL #TDOC $229.6 x 20 ; Profit: +112$',
            'data': []
        },
        {
            'msg': '✅ SELL #LAZR (на спб бирже нет) $31.88 x 130 ; Profit: +121$',
            'data': []
        },
        {
            'msg': '✅ SELL #SNAP $51.05 x 60 ; Profit: +75$',
            'data': []
        }
    ]

def test_parse_text(example_pattern, parse_text_cases):
    for case in parse_text_cases:
        data = parse_text(example_pattern, case['msg'])
        assert data == case['data']

def test_parse_text_ignore(example_pattern, parse_text_cases_ignore):
    for case in parse_text_cases_ignore:
        data = parse_text(example_pattern, case['msg'])
        assert data == case['data']
