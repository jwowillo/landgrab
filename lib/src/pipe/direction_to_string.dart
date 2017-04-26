import 'package:angular2/angular2.dart';

import 'package:landgrab/landgrab.dart';

@Pipe(name: 'directionToString')
class DirectionToStringPipe extends PipeTransform {
  String transform(Direction d) => directionToString(d);
}
