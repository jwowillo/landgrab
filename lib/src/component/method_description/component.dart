import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'method-description',
  templateUrl: 'template.html',
  pipes: const [PrettyJSONPipe],
)
class MethodDescriptionComponent {
  @Input()
  MethodDescription description;
}
