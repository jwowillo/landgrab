import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'resource',
  templateUrl: 'template.html',
  pipes: const [PrettyJSONPipe],
)
class ResourceComponent {
  @Input()
  Resource resource;
}
