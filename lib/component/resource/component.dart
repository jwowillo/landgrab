import 'package:angular2/core.dart';

import 'package:landgrab/model/api.dart';
import 'package:landgrab/pipe/pretty_json.dart';

@Component(
  selector: 'resource',
  templateUrl: 'template.html',
  pipes: const [PrettyJSONPipe],
)
class ResourceComponent {
  @Input()
  Resource resource;
}
