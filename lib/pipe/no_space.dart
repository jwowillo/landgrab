import 'package:angular2/angular2.dart';

@Pipe(name: 'noSpace')
class NoSpacePipe extends PipeTransform {
  String transform(String v) => v.replaceAll(new RegExp(r' '), '');
}
