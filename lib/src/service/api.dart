import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular2/core.dart';

import 'package:cookie/cookie.dart' as cookie;

String encode(Map<String, dynamic> map) {
  return Uri.encodeQueryComponent(new JsonEncoder().convert(map));
}

@Injectable()
class APIService {
  /// _URL is a pattern string that gets replaced by the API URL by the server.
  ///
  /// The pattern is contained here so that the Services aren't exposed to the
  /// mechanism of pattern substitution.
  static final String _URL = '{{ api }}';

  static final String _KEY = 'landgrab_token';

  String get _token => cookie.get(_KEY);

  set _token(String x) => cookie.set(_KEY, x, path: '/');

  Future<Map<String, dynamic>> request(String path,
      {Map<String, String> query}) async {
    if (_token == null) {
      await _fetchToken();
    }
    HttpRequest result;
    HttpRequest backup;
    result = await _rawRequest(path, query: query).catchError((Error e) async {
      cookie.remove('landgrab_token', path: '/');
      await _fetchToken();
      backup = await _rawRequest(path, query: query);
    });
    if (result == null) {
      result = backup;
    }
    return JSON.decode(result.responseText);
  }

  Future _fetchToken() async {
    HttpRequest request = await _rawRequest('/token').catchError((Error e) {
      throw new StateError('bad token');
    });
    Map<String, dynamic> json = JSON.decode(request.responseText);
    _token = json['data']['token'];
  }

  Future<HttpRequest> _rawRequest(String path,
      {Map<String, String> query}) async {
    String queryStr = new Uri(queryParameters: query).toString();
    return await HttpRequest.request(_URL + path + queryStr,
        method: 'GET', requestHeaders: {'Authorization': _token});
  }
}
