/// Config contains items like URLs for use in the rest of the Services.
class Config {
  /// API_URL is a pattern string that gets replaced by the API URL by the
  /// server.
  ///
  /// The pattern is contained here so that the Services aren't exposed to the
  /// mechanism of pattern substitution.
  static final String API_URL = '{{ api }}';
}
