import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'move-choice-form',
  templateUrl: 'template.html',
  styles: const ['styles.css'],
  pipes: const [DirectionToStringPipe],
)
class MoveChoiceFormComponent {
  @Input()
  int id;

  @Input()
  List<Move> moves = [];

  @Output()
  final EventEmitter<Move> changed = new EventEmitter();

  Move _checked;

  bool isChecked(Move move) => _checked != null && _checked == move;

  void emit(Move move) {
    _checked = move;
    changed.emit(move);
  }
}
