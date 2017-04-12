import 'dart:async';
import 'dart:convert';
import 'dart:html';

/// API_URL is a pattern string that gets replaced by the API URL by the server.
///
/// The pattern is contained here so that the Services aren't exposed to the
/// mechanism of pattern substitution.
final String _API_URL = '{{ api }}';

/// _api makes a request to the path at the API server with the given query
/// string.
Future<Map<String, dynamic>> api(String path,
    {Map<String, String> query}) async {
  String queryStr = new Uri(queryParameters: query).toString();
  String raw = await HttpRequest.getString(_API_URL + path + queryStr);
  return JSON.decode(raw);
}
