import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular2/core.dart';

/// API_URL is a pattern string that gets replaced by the API URL by the server.
///
/// The pattern is contained here so that the Services aren't exposed to the
/// mechanism of pattern substitution.
final String _API_URL = '{{ api }}';

String encode(Map<String, dynamic> map) {
  return Uri.encodeQueryComponent(new JsonEncoder().convert(map));
}

String _token = '';

Future _fetchToken() async {
  HttpRequest r =
      await HttpRequest.request(_API_URL + '/token').catchError((Error e) {
    throw new StateError('bad token');
  });
  Map<String, dynamic> json = JSON.decode(r.responseText);
  _token = json['data']['token'];
}

/// _api makes a request to the path at the API server with the given query
/// string.
Future<Map<String, dynamic>> api(String path,
    {Map<String, String> query}) async {
  if (_token == '') {
    await _fetchToken();
  }
  String queryStr = new Uri(queryParameters: query).toString();
  HttpRequest result;
  HttpRequest backup;
  result = await HttpRequest
      .request(_API_URL + path + queryStr, method: 'GET', requestHeaders: {
    'Authorization': _token,
  }).catchError((Error e) async {
    _token = '';
    await _fetchToken();
    backup = await HttpRequest.request(_API_URL + path + queryStr,
        method: 'GET', requestHeaders: {'Authorization': _token});
  });
  if (result == null) {
    result = backup;
  }
  return JSON.decode(result.responseText);
}
