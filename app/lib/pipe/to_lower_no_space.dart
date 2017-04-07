import 'package:angular2/angular2.dart';

@Pipe(name: 'toLowerNoSpace')
class ToLowerNoSpace extends PipeTransform {
  String transform(String v) =>
      v.toLowerCase().replaceAll(new RegExp(r' '), '');
}
