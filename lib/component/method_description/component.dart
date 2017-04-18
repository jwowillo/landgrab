import 'package:angular2/core.dart';

import 'package:landgrab/model/api.dart';
import 'package:landgrab/pipe/pretty_json.dart';

@Component(
  selector: 'method-description',
  templateUrl: 'template.html',
  pipes: const [PrettyJSONPipe],
)
class MethodDescriptionComponent {
  @Input()
  MethodDescription description;
}
