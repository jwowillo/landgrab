import 'dart:convert';

import 'package:angular2/angular2.dart';

@Pipe(name: 'prettyJSON')
class PrettyJSONPipe extends PipeTransform {
  String transform(Object o) => new JsonEncoder.withIndent('  ').convert(o);
}
