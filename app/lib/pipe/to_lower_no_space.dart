import 'package:angular2/angular2.dart';

/// ToLowerNoSpace is a pipe which accepts a String and transforms it to
/// lower-case and removes all spaces.
@Pipe(name: 'toLowerNoSpace')
class ToLowerNoSpace extends PipeTransform {
  /// transform the String to lower-case and remove all spaces.
  String transform(String v) =>
      v.toLowerCase().replaceAll(new RegExp(r' '), '');
}
