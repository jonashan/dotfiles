if type "xrandr"; then
  for m in $(xrandr --query | grep " connected" | cut -d" " -f1); do
    MONITOR=$m polybar --reload top & # htop &
  done
else
  polybar --reload top & # toph &
fi
